package bankservice

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/ivanmatyash/bank-golang/api"
	"github.com/ivanmatyash/bank-golang/sqlstore"
)

func (s *bankServer) ListTransaction(ctx context.Context, in *api.RequestTime) (*api.ResponseTransaction, error) {
	transactions := []*api.Transaction{}
	if in.End == 0 {
		in.End = math.MaxInt64
	}
	err := sqlstore.Db.Select(&transactions, "SELECT * FROM transactions WHERE timestamp >= $1 AND timestamp <= $2", in.Start, in.End)
	if err != nil {
		log.Println(err)
		return &api.ResponseTransaction{[]*api.Transaction{}}, err
	}
	log.Printf("Transactions were readed (start = %d, end = %d).", in.Start, in.End)
	return &api.ResponseTransaction{transactions}, nil
}

func (s *bankServer) StartTransaction(comment string) (*api.Transaction, error) {
	transaction := api.Transaction{}
	transaction.DiffMoney = make(map[int32]int64)
	transaction.Comment = comment
	if err := transaction.Validate(); err != nil {
		log.Println(err)
		return nil, err
	}
	sqlstore.Mutex.Lock()
	log.Printf("Start transaction '%s'.\n", comment)

	return &transaction, nil
}

func (s *bankServer) EndTransaction(transaction *api.Transaction, ok bool) error {
	defer sqlstore.Mutex.Unlock()
	transaction.Success = ok
	var nextId int32
	err := sqlstore.Db.QueryRow("select nextval ('transactions_id_seq')").Scan(&nextId)
	if err != nil {
		log.Println(err)
		return err
	}
	transaction.Id = nextId
	transaction.Timestamp = time.Now().Unix()

	query := `
		INSERT INTO transactions(
			id,
			comment, 
			success,
			timestamp
		) VALUES(
			:id,
			:comment, 
			:success,
			:timestamp
		)`

	res, err := sqlstore.Db.NamedQuery(query, transaction)
	if err != nil {
		log.Println(err)
		res.Close()
		return err
	}
	res.Close()
	saveMoneyChangesInDB(transaction)

	log.Printf("End transaction '%s'.\n", transaction.Comment)

	return nil
}

func saveMoneyChangesInDB(t *api.Transaction) error {
	type record struct {
		TransactionId int32 `db:"transaction_id"`
		AccountId     int32 `db:"account_id"`
		Diff          int64 `db:"diff"`
	}

	records := make([]record, 0)

	for k, v := range t.DiffMoney {
		records = append(records, record{t.Id, int32(k), v})
	}
	query := `
		INSERT INTO money_changes(
			transaction_id,
			account_id, 
			diff
		) VALUES(
			:transaction_id,
			:account_id, 
			:diff
		)`
	for _, record := range records {
		res, err := sqlstore.Db.NamedQuery(query, record)
		if err != nil {
			log.Println(err)
			return err
		}
		res.Close()
	}
	log.Println("Accounts changed: ", t.DiffMoney)
	return nil
}

func (s *bankServer) CancelTransaction(ctx context.Context, in *api.RequestTransaction) (*api.ResponseTransaction, error) {
	transaction, err := s.StartTransaction("Canceling transaction - " + in.Req.String())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	type record struct {
		TransactionId int32 `db:"transaction_id"`
		AccountId     int32 `db:"account_id"`
		Diff          int64
	}

	diffMoney := []*record{}

	err = sqlstore.Db.Select(&diffMoney, "SELECT * FROM money_changes WHERE transaction_id = $1", in.Req.Id)
	if err != nil {
		log.Println(err)
		if err := s.EndTransaction(transaction, false); err != nil {
			log.Println(err)
			return nil, err
		}
		return nil, err
	}

	accounts := []api.Account{}

	for _, r := range diffMoney {
		account := api.Account{}
		err := sqlstore.Db.Get(&account, "SELECT * FROM accounts WHERE id=$1", r.AccountId)
		if err != nil {
			log.Println(err)
			if err := s.EndTransaction(transaction, false); err != nil {
				log.Println(err)
				return nil, err
			}
			return nil, err
		}
		account.Balance -= r.Diff
		if err := account.Validate(); err != nil {
			log.Println(err)
			if err := s.EndTransaction(transaction, false); err != nil {
				log.Println(err)
				return nil, err
			}
			return nil, err
		}
		accounts = append(accounts, account)

		transaction.DiffMoney[r.AccountId] = -1 * r.Diff
	}

	for _, obj := range accounts {
		if _, err := s.UpdateAccount(ctx, &api.RequestAccount{&obj, obj.Id}); err != nil {
			log.Println(err)
			if err := s.EndTransaction(transaction, false); err != nil {
				log.Println(err)
				return nil, err
			}
			return nil, err
		}
	}

	if err := s.EndTransaction(transaction, true); err != nil {
		log.Println(err)
		return nil, err
	}
	return &api.ResponseTransaction{[]*api.Transaction{in.Req}}, nil
}
