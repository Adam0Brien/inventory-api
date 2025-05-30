// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: kessel/inventory/v1beta1/authz/check.proto

package authz

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
	KesselCheckService_Check_FullMethodName          = "/kessel.inventory.v1beta1.authz.KesselCheckService/Check"
	KesselCheckService_CheckForUpdate_FullMethodName = "/kessel.inventory.v1beta1.authz.KesselCheckService/CheckForUpdate"
)

// KesselCheckServiceClient is the client API for KesselCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KesselCheckServiceClient interface {
	// Checks for the existence of a single Relationship
	// (a Relation between a Resource and a Subject or Subject Set).
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	CheckForUpdate(ctx context.Context, in *CheckForUpdateRequest, opts ...grpc.CallOption) (*CheckForUpdateResponse, error)
}

type kesselCheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKesselCheckServiceClient(cc grpc.ClientConnInterface) KesselCheckServiceClient {
	return &kesselCheckServiceClient{cc}
}

func (c *kesselCheckServiceClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, KesselCheckService_Check_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kesselCheckServiceClient) CheckForUpdate(ctx context.Context, in *CheckForUpdateRequest, opts ...grpc.CallOption) (*CheckForUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckForUpdateResponse)
	err := c.cc.Invoke(ctx, KesselCheckService_CheckForUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KesselCheckServiceServer is the server API for KesselCheckService service.
// All implementations must embed UnimplementedKesselCheckServiceServer
// for forward compatibility.
type KesselCheckServiceServer interface {
	// Checks for the existence of a single Relationship
	// (a Relation between a Resource and a Subject or Subject Set).
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	CheckForUpdate(context.Context, *CheckForUpdateRequest) (*CheckForUpdateResponse, error)
	mustEmbedUnimplementedKesselCheckServiceServer()
}

// UnimplementedKesselCheckServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKesselCheckServiceServer struct{}

func (UnimplementedKesselCheckServiceServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedKesselCheckServiceServer) CheckForUpdate(context.Context, *CheckForUpdateRequest) (*CheckForUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckForUpdate not implemented")
}
func (UnimplementedKesselCheckServiceServer) mustEmbedUnimplementedKesselCheckServiceServer() {}
func (UnimplementedKesselCheckServiceServer) testEmbeddedByValue()                            {}

// UnsafeKesselCheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KesselCheckServiceServer will
// result in compilation errors.
type UnsafeKesselCheckServiceServer interface {
	mustEmbedUnimplementedKesselCheckServiceServer()
}

func RegisterKesselCheckServiceServer(s grpc.ServiceRegistrar, srv KesselCheckServiceServer) {
	// If the following call pancis, it indicates UnimplementedKesselCheckServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KesselCheckService_ServiceDesc, srv)
}

func _KesselCheckService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KesselCheckServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KesselCheckService_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KesselCheckServiceServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KesselCheckService_CheckForUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckForUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KesselCheckServiceServer).CheckForUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KesselCheckService_CheckForUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KesselCheckServiceServer).CheckForUpdate(ctx, req.(*CheckForUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KesselCheckService_ServiceDesc is the grpc.ServiceDesc for KesselCheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KesselCheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kessel.inventory.v1beta1.authz.KesselCheckService",
	HandlerType: (*KesselCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _KesselCheckService_Check_Handler,
		},
		{
			MethodName: "CheckForUpdate",
			Handler:    _KesselCheckService_CheckForUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kessel/inventory/v1beta1/authz/check.proto",
}
