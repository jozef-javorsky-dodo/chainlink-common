// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: keyvalue_store.proto

package pb

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
	KeyValueStore_StoreKeyValue_FullMethodName  = "/loop.KeyValueStore/StoreKeyValue"
	KeyValueStore_GetValueForKey_FullMethodName = "/loop.KeyValueStore/GetValueForKey"
)

// KeyValueStoreClient is the client API for KeyValueStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyValueStoreClient interface {
	StoreKeyValue(ctx context.Context, in *StoreKeyValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetValueForKey(ctx context.Context, in *GetValueForKeyRequest, opts ...grpc.CallOption) (*GetValueForKeyResponse, error)
}

type keyValueStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueStoreClient(cc grpc.ClientConnInterface) KeyValueStoreClient {
	return &keyValueStoreClient{cc}
}

func (c *keyValueStoreClient) StoreKeyValue(ctx context.Context, in *StoreKeyValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KeyValueStore_StoreKeyValue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueStoreClient) GetValueForKey(ctx context.Context, in *GetValueForKeyRequest, opts ...grpc.CallOption) (*GetValueForKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetValueForKeyResponse)
	err := c.cc.Invoke(ctx, KeyValueStore_GetValueForKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueStoreServer is the server API for KeyValueStore service.
// All implementations must embed UnimplementedKeyValueStoreServer
// for forward compatibility.
type KeyValueStoreServer interface {
	StoreKeyValue(context.Context, *StoreKeyValueRequest) (*emptypb.Empty, error)
	GetValueForKey(context.Context, *GetValueForKeyRequest) (*GetValueForKeyResponse, error)
	mustEmbedUnimplementedKeyValueStoreServer()
}

// UnimplementedKeyValueStoreServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKeyValueStoreServer struct{}

func (UnimplementedKeyValueStoreServer) StoreKeyValue(context.Context, *StoreKeyValueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreKeyValue not implemented")
}
func (UnimplementedKeyValueStoreServer) GetValueForKey(context.Context, *GetValueForKeyRequest) (*GetValueForKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValueForKey not implemented")
}
func (UnimplementedKeyValueStoreServer) mustEmbedUnimplementedKeyValueStoreServer() {}
func (UnimplementedKeyValueStoreServer) testEmbeddedByValue()                       {}

// UnsafeKeyValueStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyValueStoreServer will
// result in compilation errors.
type UnsafeKeyValueStoreServer interface {
	mustEmbedUnimplementedKeyValueStoreServer()
}

func RegisterKeyValueStoreServer(s grpc.ServiceRegistrar, srv KeyValueStoreServer) {
	// If the following call pancis, it indicates UnimplementedKeyValueStoreServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KeyValueStore_ServiceDesc, srv)
}

func _KeyValueStore_StoreKeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreKeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).StoreKeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueStore_StoreKeyValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).StoreKeyValue(ctx, req.(*StoreKeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueStore_GetValueForKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueForKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).GetValueForKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueStore_GetValueForKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).GetValueForKey(ctx, req.(*GetValueForKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyValueStore_ServiceDesc is the grpc.ServiceDesc for KeyValueStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyValueStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.KeyValueStore",
	HandlerType: (*KeyValueStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreKeyValue",
			Handler:    _KeyValueStore_StoreKeyValue_Handler,
		},
		{
			MethodName: "GetValueForKey",
			Handler:    _KeyValueStore_GetValueForKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keyvalue_store.proto",
}
