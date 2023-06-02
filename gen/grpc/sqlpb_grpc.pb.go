// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: grpc/sqlpb.proto

package grpc

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

const (
	SQLService_GetUsers_FullMethodName              = "/sqlInjection.SQLService/getUsers"
	SQLService_GetUsersWithSqlInject_FullMethodName = "/sqlInjection.SQLService/getUsersWithSqlInject"
)

// SQLServiceClient is the client API for SQLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SQLServiceClient interface {
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetUsersWithSqlInject(ctx context.Context, in *GetUsersWithSqlInjectRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
}

type sQLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSQLServiceClient(cc grpc.ClientConnInterface) SQLServiceClient {
	return &sQLServiceClient{cc}
}

func (c *sQLServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, SQLService_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sQLServiceClient) GetUsersWithSqlInject(ctx context.Context, in *GetUsersWithSqlInjectRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, SQLService_GetUsersWithSqlInject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SQLServiceServer is the server API for SQLService service.
// All implementations must embed UnimplementedSQLServiceServer
// for forward compatibility
type SQLServiceServer interface {
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	GetUsersWithSqlInject(context.Context, *GetUsersWithSqlInjectRequest) (*GetUsersResponse, error)
	mustEmbedUnimplementedSQLServiceServer()
}

// UnimplementedSQLServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSQLServiceServer struct {
}

func (UnimplementedSQLServiceServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedSQLServiceServer) GetUsersWithSqlInject(context.Context, *GetUsersWithSqlInjectRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersWithSqlInject not implemented")
}
func (UnimplementedSQLServiceServer) mustEmbedUnimplementedSQLServiceServer() {}

// UnsafeSQLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SQLServiceServer will
// result in compilation errors.
type UnsafeSQLServiceServer interface {
	mustEmbedUnimplementedSQLServiceServer()
}

func RegisterSQLServiceServer(s grpc.ServiceRegistrar, srv SQLServiceServer) {
	s.RegisterService(&SQLService_ServiceDesc, srv)
}

func _SQLService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLService_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SQLService_GetUsersWithSqlInject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersWithSqlInjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SQLServiceServer).GetUsersWithSqlInject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SQLService_GetUsersWithSqlInject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SQLServiceServer).GetUsersWithSqlInject(ctx, req.(*GetUsersWithSqlInjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SQLService_ServiceDesc is the grpc.ServiceDesc for SQLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SQLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sqlInjection.SQLService",
	HandlerType: (*SQLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUsers",
			Handler:    _SQLService_GetUsers_Handler,
		},
		{
			MethodName: "getUsersWithSqlInject",
			Handler:    _SQLService_GetUsersWithSqlInject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/sqlpb.proto",
}