// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.2
// source: eth.proto

package protoeth

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

// RpcApiClient is the client API for RpcApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcApiClient interface {
	GetBalance(ctx context.Context, in *GetBalanceReq, opts ...grpc.CallOption) (*GetBalanceResp, error)
	GetBlockNumber(ctx context.Context, in *GetBlockNumberReq, opts ...grpc.CallOption) (*GetBlockNumberResp, error)
	NewFilter(ctx context.Context, in *NewFilterReq, opts ...grpc.CallOption) (*NewFilterResp, error)
}

type rpcApiClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcApiClient(cc grpc.ClientConnInterface) RpcApiClient {
	return &rpcApiClient{cc}
}

func (c *rpcApiClient) GetBalance(ctx context.Context, in *GetBalanceReq, opts ...grpc.CallOption) (*GetBalanceResp, error) {
	out := new(GetBalanceResp)
	err := c.cc.Invoke(ctx, "/protoeth.RpcApi/getBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcApiClient) GetBlockNumber(ctx context.Context, in *GetBlockNumberReq, opts ...grpc.CallOption) (*GetBlockNumberResp, error) {
	out := new(GetBlockNumberResp)
	err := c.cc.Invoke(ctx, "/protoeth.RpcApi/getBlockNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcApiClient) NewFilter(ctx context.Context, in *NewFilterReq, opts ...grpc.CallOption) (*NewFilterResp, error) {
	out := new(NewFilterResp)
	err := c.cc.Invoke(ctx, "/protoeth.RpcApi/newFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcApiServer is the server API for RpcApi service.
// All implementations must embed UnimplementedRpcApiServer
// for forward compatibility
type RpcApiServer interface {
	GetBalance(context.Context, *GetBalanceReq) (*GetBalanceResp, error)
	GetBlockNumber(context.Context, *GetBlockNumberReq) (*GetBlockNumberResp, error)
	NewFilter(context.Context, *NewFilterReq) (*NewFilterResp, error)
	mustEmbedUnimplementedRpcApiServer()
}

// UnimplementedRpcApiServer must be embedded to have forward compatible implementations.
type UnimplementedRpcApiServer struct {
}

func (UnimplementedRpcApiServer) GetBalance(context.Context, *GetBalanceReq) (*GetBalanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedRpcApiServer) GetBlockNumber(context.Context, *GetBlockNumberReq) (*GetBlockNumberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockNumber not implemented")
}
func (UnimplementedRpcApiServer) NewFilter(context.Context, *NewFilterReq) (*NewFilterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewFilter not implemented")
}
func (UnimplementedRpcApiServer) mustEmbedUnimplementedRpcApiServer() {}

// UnsafeRpcApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcApiServer will
// result in compilation errors.
type UnsafeRpcApiServer interface {
	mustEmbedUnimplementedRpcApiServer()
}

func RegisterRpcApiServer(s grpc.ServiceRegistrar, srv RpcApiServer) {
	s.RegisterService(&RpcApi_ServiceDesc, srv)
}

func _RpcApi_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcApiServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoeth.RpcApi/getBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcApiServer).GetBalance(ctx, req.(*GetBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcApi_GetBlockNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockNumberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcApiServer).GetBlockNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoeth.RpcApi/getBlockNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcApiServer).GetBlockNumber(ctx, req.(*GetBlockNumberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcApi_NewFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewFilterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcApiServer).NewFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoeth.RpcApi/newFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcApiServer).NewFilter(ctx, req.(*NewFilterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcApi_ServiceDesc is the grpc.ServiceDesc for RpcApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoeth.RpcApi",
	HandlerType: (*RpcApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getBalance",
			Handler:    _RpcApi_GetBalance_Handler,
		},
		{
			MethodName: "getBlockNumber",
			Handler:    _RpcApi_GetBlockNumber_Handler,
		},
		{
			MethodName: "newFilter",
			Handler:    _RpcApi_NewFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eth.proto",
}
