// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: api/storage/storage.proto

package storage

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
	StorageService_UploadChunks_FullMethodName = "/storage.StorageService/UploadChunks"
	StorageService_GetChunk_FullMethodName     = "/storage.StorageService/GetChunk"
)

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	UploadChunks(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadChunksRequest, UploadChunksResponse], error)
	GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetChunkResponse], error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) UploadChunks(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadChunksRequest, UploadChunksResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StorageService_ServiceDesc.Streams[0], StorageService_UploadChunks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadChunksRequest, UploadChunksResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StorageService_UploadChunksClient = grpc.ClientStreamingClient[UploadChunksRequest, UploadChunksResponse]

func (c *storageServiceClient) GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetChunkResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StorageService_ServiceDesc.Streams[1], StorageService_GetChunk_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetChunkRequest, GetChunkResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StorageService_GetChunkClient = grpc.ServerStreamingClient[GetChunkResponse]

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility.
type StorageServiceServer interface {
	UploadChunks(grpc.ClientStreamingServer[UploadChunksRequest, UploadChunksResponse]) error
	GetChunk(*GetChunkRequest, grpc.ServerStreamingServer[GetChunkResponse]) error
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStorageServiceServer struct{}

func (UnimplementedStorageServiceServer) UploadChunks(grpc.ClientStreamingServer[UploadChunksRequest, UploadChunksResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadChunks not implemented")
}
func (UnimplementedStorageServiceServer) GetChunk(*GetChunkRequest, grpc.ServerStreamingServer[GetChunkResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetChunk not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}
func (UnimplementedStorageServiceServer) testEmbeddedByValue()                        {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	// If the following call pancis, it indicates UnimplementedStorageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_UploadChunks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServiceServer).UploadChunks(&grpc.GenericServerStream[UploadChunksRequest, UploadChunksResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StorageService_UploadChunksServer = grpc.ClientStreamingServer[UploadChunksRequest, UploadChunksResponse]

func _StorageService_GetChunk_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetChunkRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServiceServer).GetChunk(m, &grpc.GenericServerStream[GetChunkRequest, GetChunkResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StorageService_GetChunkServer = grpc.ServerStreamingServer[GetChunkResponse]

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadChunks",
			Handler:       _StorageService_UploadChunks_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetChunk",
			Handler:       _StorageService_GetChunk_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/storage/storage.proto",
}
