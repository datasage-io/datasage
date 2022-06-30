// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: src/proto/datasource/datasource.proto

package datasource

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

// DatasourceClient is the client API for Datasource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasourceClient interface {
	AddDatasources(ctx context.Context, in *AddDatasourceRequest, opts ...grpc.CallOption) (Datasource_AddDatasourcesClient, error)
	ListDatasources(ctx context.Context, in *ListDatasourceRequest, opts ...grpc.CallOption) (Datasource_ListDatasourcesClient, error)
	DeleteDatasources(ctx context.Context, in *DeleteDatasourceRequest, opts ...grpc.CallOption) (Datasource_DeleteDatasourcesClient, error)
}

type datasourceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasourceClient(cc grpc.ClientConnInterface) DatasourceClient {
	return &datasourceClient{cc}
}

func (c *datasourceClient) AddDatasources(ctx context.Context, in *AddDatasourceRequest, opts ...grpc.CallOption) (Datasource_AddDatasourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Datasource_ServiceDesc.Streams[0], "/datasource.Datasource/AddDatasources", opts...)
	if err != nil {
		return nil, err
	}
	x := &datasourceAddDatasourcesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Datasource_AddDatasourcesClient interface {
	Recv() (*MessageResponse, error)
	grpc.ClientStream
}

type datasourceAddDatasourcesClient struct {
	grpc.ClientStream
}

func (x *datasourceAddDatasourcesClient) Recv() (*MessageResponse, error) {
	m := new(MessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *datasourceClient) ListDatasources(ctx context.Context, in *ListDatasourceRequest, opts ...grpc.CallOption) (Datasource_ListDatasourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Datasource_ServiceDesc.Streams[1], "/datasource.Datasource/ListDatasources", opts...)
	if err != nil {
		return nil, err
	}
	x := &datasourceListDatasourcesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Datasource_ListDatasourcesClient interface {
	Recv() (*ListDatasourceResponse, error)
	grpc.ClientStream
}

type datasourceListDatasourcesClient struct {
	grpc.ClientStream
}

func (x *datasourceListDatasourcesClient) Recv() (*ListDatasourceResponse, error) {
	m := new(ListDatasourceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *datasourceClient) DeleteDatasources(ctx context.Context, in *DeleteDatasourceRequest, opts ...grpc.CallOption) (Datasource_DeleteDatasourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Datasource_ServiceDesc.Streams[2], "/datasource.Datasource/DeleteDatasources", opts...)
	if err != nil {
		return nil, err
	}
	x := &datasourceDeleteDatasourcesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Datasource_DeleteDatasourcesClient interface {
	Recv() (*MessageResponse, error)
	grpc.ClientStream
}

type datasourceDeleteDatasourcesClient struct {
	grpc.ClientStream
}

func (x *datasourceDeleteDatasourcesClient) Recv() (*MessageResponse, error) {
	m := new(MessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DatasourceServer is the server API for Datasource service.
// All implementations must embed UnimplementedDatasourceServer
// for forward compatibility
type DatasourceServer interface {
	AddDatasources(*AddDatasourceRequest, Datasource_AddDatasourcesServer) error
	ListDatasources(*ListDatasourceRequest, Datasource_ListDatasourcesServer) error
	DeleteDatasources(*DeleteDatasourceRequest, Datasource_DeleteDatasourcesServer) error
	mustEmbedUnimplementedDatasourceServer()
}

// UnimplementedDatasourceServer must be embedded to have forward compatible implementations.
type UnimplementedDatasourceServer struct {
}

func (UnimplementedDatasourceServer) AddDatasources(*AddDatasourceRequest, Datasource_AddDatasourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method AddDatasources not implemented")
}
func (UnimplementedDatasourceServer) ListDatasources(*ListDatasourceRequest, Datasource_ListDatasourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListDatasources not implemented")
}
func (UnimplementedDatasourceServer) DeleteDatasources(*DeleteDatasourceRequest, Datasource_DeleteDatasourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteDatasources not implemented")
}
func (UnimplementedDatasourceServer) mustEmbedUnimplementedDatasourceServer() {}

// UnsafeDatasourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasourceServer will
// result in compilation errors.
type UnsafeDatasourceServer interface {
	mustEmbedUnimplementedDatasourceServer()
}

func RegisterDatasourceServer(s grpc.ServiceRegistrar, srv DatasourceServer) {
	s.RegisterService(&Datasource_ServiceDesc, srv)
}

func _Datasource_AddDatasources_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AddDatasourceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatasourceServer).AddDatasources(m, &datasourceAddDatasourcesServer{stream})
}

type Datasource_AddDatasourcesServer interface {
	Send(*MessageResponse) error
	grpc.ServerStream
}

type datasourceAddDatasourcesServer struct {
	grpc.ServerStream
}

func (x *datasourceAddDatasourcesServer) Send(m *MessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Datasource_ListDatasources_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListDatasourceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatasourceServer).ListDatasources(m, &datasourceListDatasourcesServer{stream})
}

type Datasource_ListDatasourcesServer interface {
	Send(*ListDatasourceResponse) error
	grpc.ServerStream
}

type datasourceListDatasourcesServer struct {
	grpc.ServerStream
}

func (x *datasourceListDatasourcesServer) Send(m *ListDatasourceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Datasource_DeleteDatasources_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeleteDatasourceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatasourceServer).DeleteDatasources(m, &datasourceDeleteDatasourcesServer{stream})
}

type Datasource_DeleteDatasourcesServer interface {
	Send(*MessageResponse) error
	grpc.ServerStream
}

type datasourceDeleteDatasourcesServer struct {
	grpc.ServerStream
}

func (x *datasourceDeleteDatasourcesServer) Send(m *MessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Datasource_ServiceDesc is the grpc.ServiceDesc for Datasource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Datasource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datasource.Datasource",
	HandlerType: (*DatasourceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddDatasources",
			Handler:       _Datasource_AddDatasources_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListDatasources",
			Handler:       _Datasource_ListDatasources_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DeleteDatasources",
			Handler:       _Datasource_DeleteDatasources_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/proto/datasource/datasource.proto",
}