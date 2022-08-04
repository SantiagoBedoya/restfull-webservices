// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: datafiles/transaction.proto

package datafiles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MoneyTransactionClient is the client API for MoneyTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoneyTransactionClient interface {
	MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type moneyTransactionClient struct {
	cc grpc.ClientConnInterface
}

func NewMoneyTransactionClient(cc grpc.ClientConnInterface) MoneyTransactionClient {
	return &moneyTransactionClient{cc}
}

func (c *moneyTransactionClient) MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/datafiles.MoneyTransaction/MakeTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoneyTransactionServer is the server API for MoneyTransaction service.
// All implementations should embed UnimplementedMoneyTransactionServer
// for forward compatibility
type MoneyTransactionServer interface {
	MakeTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error)
}

// UnimplementedMoneyTransactionServer should be embedded to have forward compatible implementations.
type UnimplementedMoneyTransactionServer struct {
}

func (UnimplementedMoneyTransactionServer) MakeTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeTransaction not implemented")
}

// UnsafeMoneyTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoneyTransactionServer will
// result in compilation errors.
type UnsafeMoneyTransactionServer interface {
	mustEmbedUnimplementedMoneyTransactionServer()
}

func RegisterMoneyTransactionServer(s grpc.ServiceRegistrar, srv MoneyTransactionServer) {
	s.RegisterService(&MoneyTransaction_ServiceDesc, srv)
}

func _MoneyTransaction_MakeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyTransactionServer).MakeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datafiles.MoneyTransaction/MakeTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyTransactionServer).MakeTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MoneyTransaction_ServiceDesc is the grpc.ServiceDesc for MoneyTransaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoneyTransaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datafiles.MoneyTransaction",
	HandlerType: (*MoneyTransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeTransaction",
			Handler:    _MoneyTransaction_MakeTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datafiles/transaction.proto",
}
