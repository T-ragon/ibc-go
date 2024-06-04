// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: hello.proto

package RPC

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
	Greets_SayHello_FullMethodName = "/hello.Greets/SayHello"
)

// GreetsClient is the client API for Greets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetsClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greetsClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetsClient(cc grpc.ClientConnInterface) GreetsClient {
	return &greetsClient{cc}
}

func (c *greetsClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Greets_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetsServer is the server API for Greets service.
// All implementations must embed UnimplementedGreetsServer
// for forward compatibility
type GreetsServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	MustEmbedUnimplementedGreetsServer()
}

// UnimplementedGreetsServer must be embedded to have forward compatible implementations.
type UnimplementedGreetsServer struct {
}

func (UnimplementedGreetsServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetsServer) mustEmbedUnimplementedGreetsServer() {}

// UnsafeGreetsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetsServer will
// result in compilation errors.
type UnsafeGreetsServer interface {
	mustEmbedUnimplementedGreetsServer()
}

func RegisterGreetsServer(s grpc.ServiceRegistrar, srv GreetsServer) {
	s.RegisterService(&Greets_ServiceDesc, srv)
}

func _Greets_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetsServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greets_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetsServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greets_ServiceDesc is the grpc.ServiceDesc for Greets service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greets_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Greets",
	HandlerType: (*GreetsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greets_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
