// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package genericsapiv1

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

// FooServiceClient is the client API for FooService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FooServiceClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadFooReply, error)
}

type fooServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFooServiceClient(cc grpc.ClientConnInterface) FooServiceClient {
	return &fooServiceClient{cc}
}

func (c *fooServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadFooReply, error) {
	out := new(ReadFooReply)
	err := c.cc.Invoke(ctx, "/genericsapiv1.FooService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServiceServer is the server API for FooService service.
// All implementations should embed UnimplementedFooServiceServer
// for forward compatibility
type FooServiceServer interface {
	Read(context.Context, *ReadRequest) (*ReadFooReply, error)
}

// UnimplementedFooServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFooServiceServer struct {
}

func (UnimplementedFooServiceServer) Read(context.Context, *ReadRequest) (*ReadFooReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

// UnsafeFooServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FooServiceServer will
// result in compilation errors.
type UnsafeFooServiceServer interface {
	mustEmbedUnimplementedFooServiceServer()
}

func RegisterFooServiceServer(s grpc.ServiceRegistrar, srv FooServiceServer) {
	s.RegisterService(&FooService_ServiceDesc, srv)
}

func _FooService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genericsapiv1.FooService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FooService_ServiceDesc is the grpc.ServiceDesc for FooService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FooService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genericsapiv1.FooService",
	HandlerType: (*FooServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _FooService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/genericsapi.proto",
}

// BarServiceClient is the client API for BarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BarServiceClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadBarReply, error)
}

type barServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBarServiceClient(cc grpc.ClientConnInterface) BarServiceClient {
	return &barServiceClient{cc}
}

func (c *barServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadBarReply, error) {
	out := new(ReadBarReply)
	err := c.cc.Invoke(ctx, "/genericsapiv1.BarService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BarServiceServer is the server API for BarService service.
// All implementations should embed UnimplementedBarServiceServer
// for forward compatibility
type BarServiceServer interface {
	Read(context.Context, *ReadRequest) (*ReadBarReply, error)
}

// UnimplementedBarServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBarServiceServer struct {
}

func (UnimplementedBarServiceServer) Read(context.Context, *ReadRequest) (*ReadBarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

// UnsafeBarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BarServiceServer will
// result in compilation errors.
type UnsafeBarServiceServer interface {
	mustEmbedUnimplementedBarServiceServer()
}

func RegisterBarServiceServer(s grpc.ServiceRegistrar, srv BarServiceServer) {
	s.RegisterService(&BarService_ServiceDesc, srv)
}

func _BarService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genericsapiv1.BarService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BarService_ServiceDesc is the grpc.ServiceDesc for BarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genericsapiv1.BarService",
	HandlerType: (*BarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _BarService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/genericsapi.proto",
}