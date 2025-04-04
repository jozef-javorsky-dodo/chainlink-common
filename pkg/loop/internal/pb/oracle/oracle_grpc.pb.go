// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: oracle.proto

package oracle

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Oracle_StartOracle_FullMethodName = "/loop.Oracle/StartOracle"
	Oracle_CloseOracle_FullMethodName = "/loop.Oracle/CloseOracle"
)

// OracleClient is the client API for Oracle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OracleClient interface {
	StartOracle(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CloseOracle(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type oracleClient struct {
	cc grpc.ClientConnInterface
}

func NewOracleClient(cc grpc.ClientConnInterface) OracleClient {
	return &oracleClient{cc}
}

func (c *oracleClient) StartOracle(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Oracle_StartOracle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oracleClient) CloseOracle(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Oracle_CloseOracle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OracleServer is the server API for Oracle service.
// All implementations must embed UnimplementedOracleServer
// for forward compatibility.
type OracleServer interface {
	StartOracle(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	CloseOracle(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedOracleServer()
}

// UnimplementedOracleServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOracleServer struct{}

func (UnimplementedOracleServer) StartOracle(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartOracle not implemented")
}
func (UnimplementedOracleServer) CloseOracle(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseOracle not implemented")
}
func (UnimplementedOracleServer) mustEmbedUnimplementedOracleServer() {}
func (UnimplementedOracleServer) testEmbeddedByValue()                {}

// UnsafeOracleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OracleServer will
// result in compilation errors.
type UnsafeOracleServer interface {
	mustEmbedUnimplementedOracleServer()
}

func RegisterOracleServer(s grpc.ServiceRegistrar, srv OracleServer) {
	// If the following call pancis, it indicates UnimplementedOracleServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Oracle_ServiceDesc, srv)
}

func _Oracle_StartOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleServer).StartOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oracle_StartOracle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleServer).StartOracle(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oracle_CloseOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OracleServer).CloseOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oracle_CloseOracle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OracleServer).CloseOracle(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Oracle_ServiceDesc is the grpc.ServiceDesc for Oracle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oracle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.Oracle",
	HandlerType: (*OracleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartOracle",
			Handler:    _Oracle_StartOracle_Handler,
		},
		{
			MethodName: "CloseOracle",
			Handler:    _Oracle_CloseOracle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oracle.proto",
}
