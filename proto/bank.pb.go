// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bank.proto

/*
Package bank is a generated protocol buffer package.

It is generated from these files:
	proto/bank.proto

It has these top-level messages:
	Client
	Account
	RequestById
	ResponseAccount
	ResponseClient
*/
package bank

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
	Id       int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ClientId int32 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
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

type RequestById struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *RequestById) Reset()                    { *m = RequestById{} }
func (m *RequestById) String() string            { return proto.CompactTextString(m) }
func (*RequestById) ProtoMessage()               {}
func (*RequestById) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RequestById) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ResponseAccount struct {
	Result []*Account `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
}

func (m *ResponseAccount) Reset()                    { *m = ResponseAccount{} }
func (m *ResponseAccount) String() string            { return proto.CompactTextString(m) }
func (*ResponseAccount) ProtoMessage()               {}
func (*ResponseAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*ResponseClient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ResponseClient) GetResult() []*Client {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Client)(nil), "bank.Client")
	proto.RegisterType((*Account)(nil), "bank.Account")
	proto.RegisterType((*RequestById)(nil), "bank.RequestById")
	proto.RegisterType((*ResponseAccount)(nil), "bank.ResponseAccount")
	proto.RegisterType((*ResponseClient)(nil), "bank.ResponseClient")
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
	ListAccounts(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error)
	ReadAccount(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error)
}

type bankServiceClient struct {
	cc *grpc.ClientConn
}

func NewBankServiceClient(cc *grpc.ClientConn) BankServiceClient {
	return &bankServiceClient{cc}
}

func (c *bankServiceClient) ListClients(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseClient, error) {
	out := new(ResponseClient)
	err := grpc.Invoke(ctx, "/bank.BankService/ListClients", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) ReadClient(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseClient, error) {
	out := new(ResponseClient)
	err := grpc.Invoke(ctx, "/bank.BankService/ReadClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) ListAccounts(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error) {
	out := new(ResponseAccount)
	err := grpc.Invoke(ctx, "/bank.BankService/ListAccounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) ReadAccount(ctx context.Context, in *RequestById, opts ...grpc.CallOption) (*ResponseAccount, error) {
	out := new(ResponseAccount)
	err := grpc.Invoke(ctx, "/bank.BankService/ReadAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BankService service

type BankServiceServer interface {
	ListClients(context.Context, *RequestById) (*ResponseClient, error)
	ReadClient(context.Context, *RequestById) (*ResponseClient, error)
	ListAccounts(context.Context, *RequestById) (*ResponseAccount, error)
	ReadAccount(context.Context, *RequestById) (*ResponseAccount, error)
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
		FullMethod: "/bank.BankService/ListClients",
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
		FullMethod: "/bank.BankService/ReadClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).ReadClient(ctx, req.(*RequestById))
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
		FullMethod: "/bank.BankService/ListAccounts",
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
		FullMethod: "/bank.BankService/ReadAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).ReadAccount(ctx, req.(*RequestById))
	}
	return interceptor(ctx, in, info, handler)
}

var _BankService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bank.BankService",
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
			MethodName: "ListAccounts",
			Handler:    _BankService_ListAccounts_Handler,
		},
		{
			MethodName: "ReadAccount",
			Handler:    _BankService_ReadAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bank.proto",
}

func init() { proto.RegisterFile("proto/bank.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcb, 0x4a, 0xeb, 0x40,
	0x18, 0xc7, 0x49, 0x7a, 0xff, 0xd2, 0xeb, 0xd0, 0x1e, 0x42, 0xcf, 0x39, 0x50, 0x82, 0x42, 0x57,
	0x0d, 0x54, 0x10, 0xb7, 0xd6, 0x8d, 0xc5, 0x82, 0x32, 0x6e, 0xdc, 0xc9, 0x34, 0x33, 0xd4, 0xa1,
	0xe9, 0x4c, 0x6c, 0x52, 0x41, 0xc4, 0x8d, 0xaf, 0xe0, 0x23, 0xf9, 0x08, 0xbe, 0x82, 0x0f, 0x22,
	0x99, 0x8b, 0xd4, 0xd2, 0x45, 0x77, 0xdf, 0xed, 0xff, 0xfb, 0xfe, 0xf9, 0x32, 0xd0, 0x4e, 0xd6,
	0x32, 0x93, 0xe1, 0x9c, 0x88, 0xe5, 0x48, 0x85, 0xa8, 0x98, 0xc7, 0xfd, 0x7f, 0x0b, 0x29, 0x17,
	0x31, 0x0b, 0x49, 0xc2, 0x43, 0x22, 0x84, 0xcc, 0x48, 0xc6, 0xa5, 0x48, 0xf5, 0x4c, 0x70, 0x07,
	0xe5, 0x8b, 0x98, 0x33, 0x91, 0xa1, 0x26, 0xb8, 0x9c, 0xfa, 0xce, 0xc0, 0x19, 0x96, 0xb0, 0xcb,
	0x29, 0x42, 0x50, 0x14, 0x64, 0xc5, 0x7c, 0x77, 0xe0, 0x0c, 0x6b, 0x58, 0xc5, 0xa8, 0x0b, 0x25,
	0xb6, 0x22, 0x3c, 0xf6, 0x0b, 0xaa, 0xa8, 0x93, 0xbc, 0x9a, 0x3c, 0x48, 0xc1, 0xfc, 0xa2, 0xae,
	0xaa, 0x24, 0xb8, 0x81, 0xca, 0x79, 0x14, 0xc9, 0xcd, 0x1e, 0xf4, 0x5f, 0xa8, 0x45, 0x6a, 0xe9,
	0x3d, 0xa7, 0x8a, 0x5f, 0xc2, 0x55, 0x5d, 0x98, 0x52, 0xe4, 0x43, 0x65, 0x4e, 0x62, 0x22, 0x22,
	0xa6, 0xb6, 0x14, 0xb0, 0x4d, 0x83, 0xff, 0xe0, 0x61, 0xf6, 0xb8, 0x61, 0x69, 0x36, 0x79, 0x9e,
	0xd2, 0x5d, 0x6a, 0x70, 0x06, 0x2d, 0xcc, 0xd2, 0x44, 0x8a, 0x94, 0xd9, 0xc5, 0xc7, 0x50, 0x5e,
	0xb3, 0x74, 0x13, 0x67, 0xbe, 0x33, 0x28, 0x0c, 0xbd, 0x71, 0x63, 0xa4, 0xce, 0x63, 0xda, 0xd8,
	0x34, 0x83, 0x53, 0x68, 0x5a, 0xa5, 0x39, 0xc6, 0xd1, 0x8e, 0xb0, 0xae, 0x85, 0xba, 0x6b, 0x75,
	0xe3, 0x0f, 0x17, 0xbc, 0x09, 0x11, 0xcb, 0x5b, 0xb6, 0x7e, 0xe2, 0x11, 0x43, 0x97, 0xe0, 0xcd,
	0x78, 0x9a, 0xe9, 0xa9, 0x14, 0x75, 0xb4, 0x68, 0xcb, 0x73, 0xbf, 0x6b, 0x4b, 0xdb, 0xdb, 0x82,
	0xf6, 0xdb, 0xe7, 0xd7, 0xbb, 0x0b, 0xa8, 0x1a, 0x46, 0x46, 0x3a, 0x03, 0xc0, 0x8c, 0x50, 0xe3,
	0xe6, 0x60, 0x50, 0x4f, 0x81, 0x5a, 0xa8, 0x61, 0x41, 0xe1, 0x0b, 0xa7, 0xaf, 0xe8, 0x0a, 0xea,
	0xb9, 0x2f, 0xf3, 0xd9, 0x7b, 0x8d, 0xf5, 0x7e, 0xf3, 0xcc, 0x68, 0xd0, 0x51, 0x40, 0x0f, 0xd5,
	0x42, 0x62, 0xc5, 0xd7, 0xf9, 0x5f, 0x20, 0xd4, 0x9e, 0xf8, 0x70, 0xd6, 0x1f, 0xc5, 0x6a, 0xa3,
	0xe6, 0x0f, 0x4b, 0xb9, 0x9b, 0x97, 0xd5, 0x4b, 0x3c, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0xf5, 0x8c, 0x0a, 0xc1, 0x02, 0x00, 0x00,
}
