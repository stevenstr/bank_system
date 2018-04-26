// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bank.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	proto/bank.proto

It has these top-level messages:
	Client
	Account
	Transaction
	RequestById
	RequestAccount
	RequestClient
	ResponseAccount
	ResponseClient
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Client struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone string `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Client) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Client) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Client) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type Account struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// @inject_tag: db:"client_id"
	ClientId int32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty" db:"client_id"`
	Balance  int64 `protobuf:"varint,3,opt,name=balance" json:"balance,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Account) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Account) GetClientId() int32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Account) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type Transaction struct {
	Id        int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	DiffMoney []int64 `protobuf:"varint,2,rep,packed,name=diff_money,json=diffMoney" json:"diff_money,omitempty"`
	Comment   string  `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"`
	Success   bool    `protobuf:"varint,4,opt,name=success" json:"success,omitempty"`
	Timestamp int64   `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Transaction) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Transaction) GetDiffMoney() []int64 {
	if m != nil {
		return m.DiffMoney
	}
	return nil
}

func (m *Transaction) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Transaction) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Transaction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type RequestById struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *RequestById) Reset()                    { *m = RequestById{} }
func (m *RequestById) String() string            { return proto.CompactTextString(m) }
func (*RequestById) ProtoMessage()               {}
func (*RequestById) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestById) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RequestAccount struct {
	Req *Account `protobuf:"bytes,1,opt,name=req" json:"req,omitempty"`
	Id  int32    `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *RequestAccount) Reset()                    { *m = RequestAccount{} }
func (m *RequestAccount) String() string            { return proto.CompactTextString(m) }
func (*RequestAccount) ProtoMessage()               {}
func (*RequestAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestAccount) GetReq() *Account {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *RequestAccount) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RequestClient struct {
	Req *Client `protobuf:"bytes,1,opt,name=req" json:"req,omitempty"`
}

func (m *RequestClient) Reset()                    { *m = RequestClient{} }
func (m *RequestClient) String() string            { return proto.CompactTextString(m) }
func (*RequestClient) ProtoMessage()               {}
func (*RequestClient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RequestClient) GetReq() *Client {
	if m != nil {
		return m.Req
	}
	return nil
}

type ResponseAccount struct {
	Result []*Account `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
}

func (m *ResponseAccount) Reset()                    { *m = ResponseAccount{} }
func (m *ResponseAccount) String() string            { return proto.CompactTextString(m) }
func (*ResponseAccount) ProtoMessage()               {}
func (*ResponseAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ResponseAccount) GetResult() []*Account {
	if m != nil {
		return m.Result
	}
	return nil
}

type ResponseClient struct {
	Result []*Client `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
}

func (m *ResponseClient) Reset()                    { *m = ResponseClient{} }
func (m *ResponseClient) String() string            { return proto.CompactTextString(m) }
func (*ResponseClient) ProtoMessage()               {}
func (*ResponseClient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResponseClient) GetResult() []*Client {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Client)(nil), "api.Client")
	proto.RegisterType((*Account)(nil), "api.Account")
	proto.RegisterType((*Transaction)(nil), "api.Transaction")
	proto.RegisterType((*RequestById)(nil), "api.RequestById")
	proto.RegisterType((*RequestAccount)(nil), "api.RequestAccount")
	proto.RegisterType((*RequestClient)(nil), "api.RequestClient")
	proto.RegisterType((*ResponseAccount)(nil), "api.ResponseAccount")
	proto.RegisterType((*ResponseClient)(nil), "api.ResponseClient")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BankService service

type BankServiceClient interface {
	ListClients(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseClient, error)
	ReadClient(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseClient, error)
	ListAccountsByClient(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error)
	CreateClient(ctx context.Context, in *RequestClient, opts ...grpc.CallOption) (*ResponseClient, error)
	DeleteClient(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseClient, error)
	ListAccounts(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error)
	ReadAccount(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error)
	CreateAccount(ctx context.Context, in *RequestAccount, opts ...grpc.CallOption) (*ResponseAccount, error)
	UpdateAccount(ctx context.Context, in *RequestAccount, opts ...grpc.CallOption) (*ResponseAccount, error)
	DeleteAccount(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error)
}

type bankServiceClient struct {
	cc *grpc.ClientConn
}

func NewBankServiceClient(cc *grpc.ClientConn) BankServiceClient {
	return &bankServiceClient{cc}
}

func (c *bankServiceClient) ListClients(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseClient, error) {
	out := new(ResponseClient)
	err := grpc.Invoke(ctx, "/api.BankService/ListClients", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) ReadClient(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseClient, error) {
	out := new(ResponseClient)
	err := grpc.Invoke(ctx, "/api.BankService/ReadClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) ListAccountsByClient(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error) {
	out := new(ResponseAccount)
	err := grpc.Invoke(ctx, "/api.BankService/ListAccountsByClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) CreateClient(ctx context.Context, in *RequestClient, opts ...grpc.CallOption) (*ResponseClient, error) {
	out := new(ResponseClient)
	err := grpc.Invoke(ctx, "/api.BankService/CreateClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) DeleteClient(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseClient, error) {
	out := new(ResponseClient)
	err := grpc.Invoke(ctx, "/api.BankService/DeleteClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) ListAccounts(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error) {
	out := new(ResponseAccount)
	err := grpc.Invoke(ctx, "/api.BankService/ListAccounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) ReadAccount(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error) {
	out := new(ResponseAccount)
	err := grpc.Invoke(ctx, "/api.BankService/ReadAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) CreateAccount(ctx context.Context, in *RequestAccount, opts ...grpc.CallOption) (*ResponseAccount, error) {
	out := new(ResponseAccount)
	err := grpc.Invoke(ctx, "/api.BankService/CreateAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) UpdateAccount(ctx context.Context, in *RequestAccount, opts ...grpc.CallOption) (*ResponseAccount, error) {
	out := new(ResponseAccount)
	err := grpc.Invoke(ctx, "/api.BankService/UpdateAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) DeleteAccount(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error) {
	out := new(ResponseAccount)
	err := grpc.Invoke(ctx, "/api.BankService/DeleteAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BankService service

type BankServiceServer interface {
	ListClients(context.Context, *RequestById) (*ResponseClient, error)
	ReadClient(context.Context, *RequestById) (*ResponseClient, error)
	ListAccountsByClient(context.Context, *RequestById) (*ResponseAccount, error)
	CreateClient(context.Context, *RequestClient) (*ResponseClient, error)
	DeleteClient(context.Context, *RequestById) (*ResponseClient, error)
	ListAccounts(context.Context, *RequestById) (*ResponseAccount, error)
	ReadAccount(context.Context, *RequestById) (*ResponseAccount, error)
	CreateAccount(context.Context, *RequestAccount) (*ResponseAccount, error)
	UpdateAccount(context.Context, *RequestAccount) (*ResponseAccount, error)
	DeleteAccount(context.Context, *RequestById) (*ResponseAccount, error)
}

func RegisterBankServiceServer(s *grpc.Server, srv BankServiceServer) {
	s.RegisterService(&_BankService_serviceDesc, srv)
}

func _BankService_ListClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).ListClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/ListClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).ListClients(ctx, req.(*RequestById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_ReadClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).ReadClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/ReadClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).ReadClient(ctx, req.(*RequestById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_ListAccountsByClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).ListAccountsByClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/ListAccountsByClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).ListAccountsByClient(ctx, req.(*RequestById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).CreateClient(ctx, req.(*RequestClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).DeleteClient(ctx, req.(*RequestById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).ListAccounts(ctx, req.(*RequestById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_ReadAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).ReadAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/ReadAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).ReadAccount(ctx, req.(*RequestById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).CreateAccount(ctx, req.(*RequestAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).UpdateAccount(ctx, req.(*RequestAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BankService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).DeleteAccount(ctx, req.(*RequestById))
	}
	return interceptor(ctx, in, info, handler)
}

var _BankService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.BankService",
	HandlerType: (*BankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListClients",
			Handler:    _BankService_ListClients_Handler,
		},
		{
			MethodName: "ReadClient",
			Handler:    _BankService_ReadClient_Handler,
		},
		{
			MethodName: "ListAccountsByClient",
			Handler:    _BankService_ListAccountsByClient_Handler,
		},
		{
			MethodName: "CreateClient",
			Handler:    _BankService_CreateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _BankService_DeleteClient_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _BankService_ListAccounts_Handler,
		},
		{
			MethodName: "ReadAccount",
			Handler:    _BankService_ReadAccount_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _BankService_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _BankService_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _BankService_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bank.proto",
}

func init() { proto.RegisterFile("proto/bank.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0x15, 0xbb, 0x49, 0x9b, 0x71, 0x92, 0xe6, 0xbf, 0x4d, 0x2b, 0x2b, 0x6d, 0xaa, 0xc8,
	0x7f, 0x0e, 0x51, 0x0f, 0xb1, 0x54, 0x84, 0x90, 0x38, 0x41, 0x8a, 0x90, 0x82, 0x5a, 0x40, 0x06,
	0x04, 0x27, 0xaa, 0x8d, 0x3d, 0x2d, 0xab, 0xda, 0x6b, 0xd7, 0xeb, 0x20, 0x45, 0x88, 0x0b, 0x0f,
	0xc0, 0x85, 0x87, 0xe1, 0x41, 0x78, 0x05, 0x1e, 0x04, 0x79, 0x77, 0x9d, 0x38, 0x29, 0x54, 0xa1,
	0x37, 0xef, 0x37, 0x9e, 0xdf, 0x7c, 0xdf, 0x78, 0x65, 0x68, 0x27, 0x69, 0x9c, 0xc5, 0xee, 0x84,
	0xf2, 0xab, 0xa1, 0x7c, 0x24, 0x26, 0x4d, 0x58, 0xf7, 0xe0, 0x32, 0x8e, 0x2f, 0x43, 0x74, 0x69,
	0xc2, 0x5c, 0xca, 0x79, 0x9c, 0xd1, 0x8c, 0xc5, 0x5c, 0xa8, 0x57, 0x9c, 0xf7, 0x50, 0x3b, 0x09,
	0x19, 0xf2, 0x8c, 0xb4, 0xc0, 0x60, 0x81, 0x5d, 0xe9, 0x57, 0x06, 0x55, 0xcf, 0x60, 0x01, 0x21,
	0xb0, 0xc1, 0x69, 0x84, 0xb6, 0xd1, 0xaf, 0x0c, 0xea, 0x9e, 0x7c, 0x26, 0x1d, 0xa8, 0x62, 0x44,
	0x59, 0x68, 0x9b, 0x52, 0x54, 0x87, 0x5c, 0x4d, 0x3e, 0xc6, 0x1c, 0xed, 0x0d, 0xa5, 0xca, 0x83,
	0xf3, 0x0a, 0x36, 0x9f, 0xf8, 0x7e, 0x3c, 0xfd, 0x03, 0x7a, 0x1f, 0xea, 0xbe, 0x1c, 0x7a, 0xce,
	0x02, 0xc9, 0xaf, 0x7a, 0x5b, 0x4a, 0x18, 0x07, 0xc4, 0x86, 0xcd, 0x09, 0x0d, 0x29, 0xf7, 0x51,
	0x4e, 0x31, 0xbd, 0xe2, 0xe8, 0x7c, 0xab, 0x80, 0xf5, 0x26, 0xa5, 0x5c, 0x50, 0x3f, 0x8f, 0x70,
	0x03, 0xdb, 0x03, 0x08, 0xd8, 0xc5, 0xc5, 0x79, 0x14, 0x73, 0x9c, 0xd9, 0x46, 0xdf, 0x1c, 0x98,
	0x5e, 0x3d, 0x57, 0xce, 0x72, 0x21, 0x07, 0xfb, 0x71, 0x14, 0x21, 0xcf, 0xb4, 0xfd, 0xe2, 0x98,
	0x57, 0xc4, 0xd4, 0xf7, 0x51, 0x08, 0x19, 0x61, 0xcb, 0x2b, 0x8e, 0xe4, 0x00, 0xea, 0x19, 0x8b,
	0x50, 0x64, 0x34, 0x4a, 0xec, 0xaa, 0xb4, 0xb3, 0x10, 0x9c, 0x1e, 0x58, 0x1e, 0x5e, 0x4f, 0x51,
	0x64, 0xa3, 0xd9, 0x38, 0x58, 0xf5, 0xe3, 0x3c, 0x86, 0x96, 0x2e, 0x17, 0x8b, 0x38, 0x04, 0x33,
	0xc5, 0x6b, 0xf9, 0x8a, 0x75, 0xdc, 0x18, 0xd2, 0x84, 0x0d, 0x75, 0xc9, 0xcb, 0x0b, 0x9a, 0x60,
	0xcc, 0x09, 0x43, 0x68, 0x6a, 0x82, 0xfe, 0x48, 0xbd, 0x32, 0xc0, 0x92, 0x00, 0x55, 0x91, 0xfd,
	0xce, 0x43, 0xd8, 0xf6, 0x50, 0x24, 0x31, 0x17, 0x58, 0x8c, 0xbc, 0x07, 0xb5, 0x14, 0xc5, 0x34,
	0xcc, 0xec, 0x4a, 0xdf, 0xbc, 0x31, 0x55, 0xd7, 0x9c, 0x07, 0xb9, 0x55, 0xd5, 0xa8, 0x27, 0xfd,
	0xbf, 0xd2, 0xb7, 0x34, 0x4c, 0x97, 0x8e, 0x7f, 0xd4, 0xc0, 0x1a, 0x51, 0x7e, 0xf5, 0x1a, 0xd3,
	0x4f, 0xcc, 0x47, 0xf2, 0x0c, 0xac, 0x53, 0x56, 0x98, 0x15, 0xa4, 0x2d, 0x7b, 0x4a, 0x2b, 0xea,
	0xee, 0x68, 0xa5, 0x3c, 0xca, 0x69, 0x7f, 0xfd, 0xf9, 0xeb, 0xbb, 0x01, 0x64, 0xcb, 0xf5, 0x75,
	0xe3, 0x73, 0x00, 0x0f, 0x69, 0xa0, 0xad, 0xac, 0x89, 0xd9, 0x95, 0x98, 0x6d, 0xd2, 0x2c, 0x30,
	0xee, 0x67, 0x16, 0x7c, 0x21, 0x1f, 0xa0, 0x93, 0x7b, 0xd2, 0x89, 0xc5, 0x68, 0xf6, 0x57, 0x6a,
	0x67, 0x89, 0xaa, 0x1b, 0x9c, 0x43, 0x89, 0xb5, 0xc9, 0xde, 0x12, 0xd6, 0xa5, 0x9a, 0x47, 0x5e,
	0x40, 0xe3, 0x24, 0x45, 0x9a, 0x15, 0x8b, 0x23, 0x65, 0xae, 0xd2, 0x6e, 0xf5, 0xeb, 0xcc, 0x63,
	0x3f, 0x92, 0x77, 0xe0, 0x14, 0x1a, 0x4f, 0x31, 0xc4, 0x39, 0xef, 0xdf, 0xd2, 0x1f, 0xad, 0xa4,
	0x1f, 0x43, 0xa3, 0x9c, 0x7e, 0xed, 0xd4, 0xff, 0x49, 0x9c, 0x45, 0xea, 0x8b, 0xa0, 0x67, 0xf9,
	0x6d, 0xa7, 0x41, 0x71, 0xb1, 0xd6, 0x25, 0xed, 0x49, 0x52, 0x9b, 0xb4, 0xe6, 0x24, 0xe5, 0xcc,
	0x83, 0xa6, 0xda, 0x5b, 0x01, 0xdc, 0x29, 0x03, 0xb5, 0x78, 0x3b, 0xd3, 0x59, 0xb8, 0x53, 0xbb,
	0x7b, 0x07, 0xcd, 0xb7, 0x49, 0x70, 0x37, 0xe6, 0xbe, 0x64, 0xee, 0x76, 0x57, 0x7c, 0x2a, 0xf0,
	0x4b, 0x68, 0xaa, 0x8f, 0x72, 0xc7, 0xf4, 0x47, 0x2b, 0xd4, 0x49, 0x4d, 0xfe, 0x7e, 0xef, 0xff,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x7d, 0x74, 0x4a, 0xb5, 0x05, 0x00, 0x00,
}
