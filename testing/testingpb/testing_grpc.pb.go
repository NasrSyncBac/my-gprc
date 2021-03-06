// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package testingpb

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestingServiceClient is the client API for TestingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestingServiceClient interface {
}

type testingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestingServiceClient(cc grpc.ClientConnInterface) TestingServiceClient {
	return &testingServiceClient{cc}
}

// TestingServiceServer is the server API for TestingService service.
// All implementations must embed UnimplementedTestingServiceServer
// for forward compatibility
type TestingServiceServer interface {
	mustEmbedUnimplementedTestingServiceServer()
}

// UnimplementedTestingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestingServiceServer struct {
}

func (UnimplementedTestingServiceServer) mustEmbedUnimplementedTestingServiceServer() {}

// UnsafeTestingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestingServiceServer will
// result in compilation errors.
type UnsafeTestingServiceServer interface {
	mustEmbedUnimplementedTestingServiceServer()
}

func RegisterTestingServiceServer(s grpc.ServiceRegistrar, srv TestingServiceServer) {
	s.RegisterService(&TestingService_ServiceDesc, srv)
}

// TestingService_ServiceDesc is the grpc.ServiceDesc for TestingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testing.TestingService",
	HandlerType: (*TestingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "testingpb/testing.proto",
}
