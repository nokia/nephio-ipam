// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pkg/alloc/allocpb/alloc.proto

package allocpb

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

// AllocationClient is the client API for Allocation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AllocationClient interface {
	Allocation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	DeAllocation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type allocationClient struct {
	cc grpc.ClientConnInterface
}

func NewAllocationClient(cc grpc.ClientConnInterface) AllocationClient {
	return &allocationClient{cc}
}

func (c *allocationClient) Allocation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/alloc.Allocation/Allocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationClient) DeAllocation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/alloc.Allocation/DeAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllocationServer is the server API for Allocation service.
// All implementations must embed UnimplementedAllocationServer
// for forward compatibility
type AllocationServer interface {
	Allocation(context.Context, *Request) (*Response, error)
	DeAllocation(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedAllocationServer()
}

// UnimplementedAllocationServer must be embedded to have forward compatible implementations.
type UnimplementedAllocationServer struct {
}

func (UnimplementedAllocationServer) Allocation(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allocation not implemented")
}
func (UnimplementedAllocationServer) DeAllocation(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeAllocation not implemented")
}
func (UnimplementedAllocationServer) mustEmbedUnimplementedAllocationServer() {}

// UnsafeAllocationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AllocationServer will
// result in compilation errors.
type UnsafeAllocationServer interface {
	mustEmbedUnimplementedAllocationServer()
}

func RegisterAllocationServer(s grpc.ServiceRegistrar, srv AllocationServer) {
	s.RegisterService(&Allocation_ServiceDesc, srv)
}

func _Allocation_Allocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServer).Allocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alloc.Allocation/Allocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServer).Allocation(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Allocation_DeAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServer).DeAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alloc.Allocation/DeAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServer).DeAllocation(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Allocation_ServiceDesc is the grpc.ServiceDesc for Allocation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Allocation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alloc.Allocation",
	HandlerType: (*AllocationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allocation",
			Handler:    _Allocation_Allocation_Handler,
		},
		{
			MethodName: "DeAllocation",
			Handler:    _Allocation_DeAllocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/alloc/allocpb/alloc.proto",
}
