// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: class.proto

package class

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

// ClassClient is the client API for Class service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClassClient interface {
	AddClass(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	ListClass(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	DeleteClass(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type classClient struct {
	cc grpc.ClientConnInterface
}

func NewClassClient(cc grpc.ClientConnInterface) ClassClient {
	return &classClient{cc}
}

func (c *classClient) AddClass(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/class.Class/AddClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classClient) ListClass(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/class.Class/ListClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classClient) DeleteClass(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/class.Class/DeleteClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassServer is the server API for Class service.
// All implementations must embed UnimplementedClassServer
// for forward compatibility
type ClassServer interface {
	AddClass(context.Context, *CreateRequest) (*MessageResponse, error)
	ListClass(context.Context, *ListRequest) (*ListResponse, error)
	DeleteClass(context.Context, *DeleteRequest) (*MessageResponse, error)
	mustEmbedUnimplementedClassServer()
}

// UnimplementedClassServer must be embedded to have forward compatible implementations.
type UnimplementedClassServer struct {
}

func (UnimplementedClassServer) AddClass(context.Context, *CreateRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClass not implemented")
}
func (UnimplementedClassServer) ListClass(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClass not implemented")
}
func (UnimplementedClassServer) DeleteClass(context.Context, *DeleteRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClass not implemented")
}
func (UnimplementedClassServer) mustEmbedUnimplementedClassServer() {}

// UnsafeClassServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClassServer will
// result in compilation errors.
type UnsafeClassServer interface {
	mustEmbedUnimplementedClassServer()
}

func RegisterClassServer(s grpc.ServiceRegistrar, srv ClassServer) {
	s.RegisterService(&Class_ServiceDesc, srv)
}

func _Class_AddClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServer).AddClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/class.Class/AddClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServer).AddClass(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Class_ListClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServer).ListClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/class.Class/ListClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServer).ListClass(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Class_DeleteClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServer).DeleteClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/class.Class/DeleteClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServer).DeleteClass(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Class_ServiceDesc is the grpc.ServiceDesc for Class service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Class_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "class.Class",
	HandlerType: (*ClassServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddClass",
			Handler:    _Class_AddClass_Handler,
		},
		{
			MethodName: "ListClass",
			Handler:    _Class_ListClass_Handler,
		},
		{
			MethodName: "DeleteClass",
			Handler:    _Class_DeleteClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "class.proto",
}
