package sqlstore

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	Db  *sqlx.DB
	err error
)

func InitDB() error {
	pgDbname := "bank"
	pgUser := "postgres"
	pgHost := "172.18.0.2"
	pgPassword := "postgres"
	pgSslmode := "disable"
	pgPort := "5432"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s"+
		"dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDbname,
		pgSslmode)

	Db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func CloseDB() error {
	err = Db.Close()
	return err
}