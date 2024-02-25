// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: cernunnos/orders/v1/orders.proto

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
	CernunnosOrdersAPI_NewOrder_FullMethodName     = "/cernunnos.orders.v1.CernunnosOrdersAPI/NewOrder"
	CernunnosOrdersAPI_GetOrders_FullMethodName    = "/cernunnos.orders.v1.CernunnosOrdersAPI/GetOrders"
	CernunnosOrdersAPI_UpdateOrder_FullMethodName  = "/cernunnos.orders.v1.CernunnosOrdersAPI/UpdateOrder"
	CernunnosOrdersAPI_DeleteOrders_FullMethodName = "/cernunnos.orders.v1.CernunnosOrdersAPI/DeleteOrders"
)

// CernunnosOrdersAPIClient is the client API for CernunnosOrdersAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CernunnosOrdersAPIClient interface {
	NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*DeleteOrdersResponse, error)
}

type cernunnosOrdersAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCernunnosOrdersAPIClient(cc grpc.ClientConnInterface) CernunnosOrdersAPIClient {
	return &cernunnosOrdersAPIClient{cc}
}

func (c *cernunnosOrdersAPIClient) NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*NewOrderResponse, error) {
	out := new(NewOrderResponse)
	err := c.cc.Invoke(ctx, CernunnosOrdersAPI_NewOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cernunnosOrdersAPIClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, CernunnosOrdersAPI_GetOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cernunnosOrdersAPIClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, CernunnosOrdersAPI_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cernunnosOrdersAPIClient) DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*DeleteOrdersResponse, error) {
	out := new(DeleteOrdersResponse)
	err := c.cc.Invoke(ctx, CernunnosOrdersAPI_DeleteOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CernunnosOrdersAPIServer is the server API for CernunnosOrdersAPI service.
// All implementations must embed UnimplementedCernunnosOrdersAPIServer
// for forward compatibility
type CernunnosOrdersAPIServer interface {
	NewOrder(context.Context, *NewOrderRequest) (*NewOrderResponse, error)
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	DeleteOrders(context.Context, *DeleteOrdersRequest) (*DeleteOrdersResponse, error)
	mustEmbedUnimplementedCernunnosOrdersAPIServer()
}

// UnimplementedCernunnosOrdersAPIServer must be embedded to have forward compatible implementations.
type UnimplementedCernunnosOrdersAPIServer struct {
}

func (UnimplementedCernunnosOrdersAPIServer) NewOrder(context.Context, *NewOrderRequest) (*NewOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (UnimplementedCernunnosOrdersAPIServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedCernunnosOrdersAPIServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedCernunnosOrdersAPIServer) DeleteOrders(context.Context, *DeleteOrdersRequest) (*DeleteOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrders not implemented")
}
func (UnimplementedCernunnosOrdersAPIServer) mustEmbedUnimplementedCernunnosOrdersAPIServer() {}

// UnsafeCernunnosOrdersAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CernunnosOrdersAPIServer will
// result in compilation errors.
type UnsafeCernunnosOrdersAPIServer interface {
	mustEmbedUnimplementedCernunnosOrdersAPIServer()
}

func RegisterCernunnosOrdersAPIServer(s grpc.ServiceRegistrar, srv CernunnosOrdersAPIServer) {
	s.RegisterService(&CernunnosOrdersAPI_ServiceDesc, srv)
}

func _CernunnosOrdersAPI_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CernunnosOrdersAPIServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CernunnosOrdersAPI_NewOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CernunnosOrdersAPIServer).NewOrder(ctx, req.(*NewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CernunnosOrdersAPI_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CernunnosOrdersAPIServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CernunnosOrdersAPI_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CernunnosOrdersAPIServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CernunnosOrdersAPI_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CernunnosOrdersAPIServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CernunnosOrdersAPI_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CernunnosOrdersAPIServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CernunnosOrdersAPI_DeleteOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CernunnosOrdersAPIServer).DeleteOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CernunnosOrdersAPI_DeleteOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CernunnosOrdersAPIServer).DeleteOrders(ctx, req.(*DeleteOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CernunnosOrdersAPI_ServiceDesc is the grpc.ServiceDesc for CernunnosOrdersAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CernunnosOrdersAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cernunnos.orders.v1.CernunnosOrdersAPI",
	HandlerType: (*CernunnosOrdersAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewOrder",
			Handler:    _CernunnosOrdersAPI_NewOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _CernunnosOrdersAPI_GetOrders_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _CernunnosOrdersAPI_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrders",
			Handler:    _CernunnosOrdersAPI_DeleteOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cernunnos/orders/v1/orders.proto",
}
