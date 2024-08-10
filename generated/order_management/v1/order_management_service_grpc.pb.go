// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: order_management/v1/order_management_service.proto

package order_management_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OrderManagementService_PlaceOrder_FullMethodName = "/github.com.mike_jacks.pizza_co.protos.order_management.v1.OrderManagementService/PlaceOrder"
)

// OrderManagementServiceClient is the client API for OrderManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementServiceClient interface {
	PlaceOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderResponse], error)
}

type orderManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementServiceClient(cc grpc.ClientConnInterface) OrderManagementServiceClient {
	return &orderManagementServiceClient{cc}
}

func (c *orderManagementServiceClient) PlaceOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderManagementService_ServiceDesc.Streams[0], OrderManagementService_PlaceOrder_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[OrderRequest, OrderResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderManagementService_PlaceOrderClient = grpc.ServerStreamingClient[OrderResponse]

// OrderManagementServiceServer is the server API for OrderManagementService service.
// All implementations must embed UnimplementedOrderManagementServiceServer
// for forward compatibility.
type OrderManagementServiceServer interface {
	PlaceOrder(*OrderRequest, grpc.ServerStreamingServer[OrderResponse]) error
	mustEmbedUnimplementedOrderManagementServiceServer()
}

// UnimplementedOrderManagementServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderManagementServiceServer struct{}

func (UnimplementedOrderManagementServiceServer) PlaceOrder(*OrderRequest, grpc.ServerStreamingServer[OrderResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedOrderManagementServiceServer) mustEmbedUnimplementedOrderManagementServiceServer() {
}
func (UnimplementedOrderManagementServiceServer) testEmbeddedByValue() {}

// UnsafeOrderManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServiceServer will
// result in compilation errors.
type UnsafeOrderManagementServiceServer interface {
	mustEmbedUnimplementedOrderManagementServiceServer()
}

func RegisterOrderManagementServiceServer(s grpc.ServiceRegistrar, srv OrderManagementServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderManagementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderManagementService_ServiceDesc, srv)
}

func _OrderManagementService_PlaceOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServiceServer).PlaceOrder(m, &grpc.GenericServerStream[OrderRequest, OrderResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OrderManagementService_PlaceOrderServer = grpc.ServerStreamingServer[OrderResponse]

// OrderManagementService_ServiceDesc is the grpc.ServiceDesc for OrderManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.mike_jacks.pizza_co.protos.order_management.v1.OrderManagementService",
	HandlerType: (*OrderManagementServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PlaceOrder",
			Handler:       _OrderManagementService_PlaceOrder_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "order_management/v1/order_management_service.proto",
}
