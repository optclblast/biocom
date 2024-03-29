// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: warden/user/v1/user.proto

package v1

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
	UsersAPI_GetUsers_FullMethodName = "/wrdn.user.v1.UsersAPI/GetUsers"
)

// UsersAPIClient is the client API for UsersAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersAPIClient interface {
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
}

type usersAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersAPIClient(cc grpc.ClientConnInterface) UsersAPIClient {
	return &usersAPIClient{cc}
}

func (c *usersAPIClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, UsersAPI_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersAPIServer is the server API for UsersAPI service.
// All implementations must embed UnimplementedUsersAPIServer
// for forward compatibility
type UsersAPIServer interface {
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	mustEmbedUnimplementedUsersAPIServer()
}

// UnimplementedUsersAPIServer must be embedded to have forward compatible implementations.
type UnimplementedUsersAPIServer struct {
}

func (UnimplementedUsersAPIServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUsersAPIServer) mustEmbedUnimplementedUsersAPIServer() {}

// UnsafeUsersAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersAPIServer will
// result in compilation errors.
type UnsafeUsersAPIServer interface {
	mustEmbedUnimplementedUsersAPIServer()
}

func RegisterUsersAPIServer(s grpc.ServiceRegistrar, srv UsersAPIServer) {
	s.RegisterService(&UsersAPI_ServiceDesc, srv)
}

func _UsersAPI_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAPIServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAPI_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAPIServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersAPI_ServiceDesc is the grpc.ServiceDesc for UsersAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wrdn.user.v1.UsersAPI",
	HandlerType: (*UsersAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _UsersAPI_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warden/user/v1/user.proto",
}
