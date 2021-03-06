// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// BibliotecaClient is the client API for Biblioteca service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BibliotecaClient interface {
	Show(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*Book, error)
}

type bibliotecaClient struct {
	cc grpc.ClientConnInterface
}

func NewBibliotecaClient(cc grpc.ClientConnInterface) BibliotecaClient {
	return &bibliotecaClient{cc}
}

func (c *bibliotecaClient) Show(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/api.Biblioteca/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BibliotecaServer is the server API for Biblioteca service.
// All implementations must embed UnimplementedBibliotecaServer
// for forward compatibility
type BibliotecaServer interface {
	Show(context.Context, *BookRequest) (*Book, error)
	mustEmbedUnimplementedBibliotecaServer()
}

// UnimplementedBibliotecaServer must be embedded to have forward compatible implementations.
type UnimplementedBibliotecaServer struct {
}

func (UnimplementedBibliotecaServer) Show(context.Context, *BookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedBibliotecaServer) mustEmbedUnimplementedBibliotecaServer() {}

// UnsafeBibliotecaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BibliotecaServer will
// result in compilation errors.
type UnsafeBibliotecaServer interface {
	mustEmbedUnimplementedBibliotecaServer()
}

func RegisterBibliotecaServer(s grpc.ServiceRegistrar, srv BibliotecaServer) {
	s.RegisterService(&Biblioteca_ServiceDesc, srv)
}

func _Biblioteca_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BibliotecaServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Biblioteca/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BibliotecaServer).Show(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Biblioteca_ServiceDesc is the grpc.ServiceDesc for Biblioteca service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Biblioteca_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Biblioteca",
	HandlerType: (*BibliotecaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Show",
			Handler:    _Biblioteca_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/biblioteca.proto",
}
