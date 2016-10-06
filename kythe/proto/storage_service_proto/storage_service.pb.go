// Code generated by protoc-gen-gogo.
// source: kythe/proto/storage_service.proto
// DO NOT EDIT!

/*
	Package storage_service_proto is a generated protocol buffer package.

	It is generated from these files:
		kythe/proto/storage_service.proto

	It has these top-level messages:
*/
package storage_service_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import kythe_proto "kythe.io/kythe/proto/storage_proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for GraphStore service

type GraphStoreClient interface {
	// Read responds with all Entry messages that match the given ReadRequest.
	// The Read operation should be implemented with time complexity proportional
	// to the size of the return set.
	Read(ctx context.Context, in *kythe_proto.ReadRequest, opts ...grpc.CallOption) (GraphStore_ReadClient, error)
	// Scan responds with all Entry messages matching the given ScanRequest.  If a
	// ScanRequest field is empty, any entry value for that field matches and will
	// be returned.  Scan is similar to Read, but with no time complexity
	// restrictions.
	Scan(ctx context.Context, in *kythe_proto.ScanRequest, opts ...grpc.CallOption) (GraphStore_ScanClient, error)
	// Write atomically inserts or updates a collection of entries into the store.
	// Each update is a tuple of the form (kind, target, fact, value).  For each
	// such update, an entry (source, kind, target, fact, value) is written into
	// the store, replacing any existing entry (source, kind, target, fact,
	// value') that may exist.  Note that this operation cannot delete any data
	// from the store; entries are only ever inserted or updated.  Apart from
	// acting atomically, no other constraints are placed on the implementation.
	Write(ctx context.Context, in *kythe_proto.WriteRequest, opts ...grpc.CallOption) (*kythe_proto.WriteReply, error)
}

type graphStoreClient struct {
	cc *grpc.ClientConn
}

func NewGraphStoreClient(cc *grpc.ClientConn) GraphStoreClient {
	return &graphStoreClient{cc}
}

func (c *graphStoreClient) Read(ctx context.Context, in *kythe_proto.ReadRequest, opts ...grpc.CallOption) (GraphStore_ReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GraphStore_serviceDesc.Streams[0], c.cc, "/kythe.proto.GraphStore/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &graphStoreReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GraphStore_ReadClient interface {
	Recv() (*kythe_proto.Entry, error)
	grpc.ClientStream
}

type graphStoreReadClient struct {
	grpc.ClientStream
}

func (x *graphStoreReadClient) Recv() (*kythe_proto.Entry, error) {
	m := new(kythe_proto.Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *graphStoreClient) Scan(ctx context.Context, in *kythe_proto.ScanRequest, opts ...grpc.CallOption) (GraphStore_ScanClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GraphStore_serviceDesc.Streams[1], c.cc, "/kythe.proto.GraphStore/Scan", opts...)
	if err != nil {
		return nil, err
	}
	x := &graphStoreScanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GraphStore_ScanClient interface {
	Recv() (*kythe_proto.Entry, error)
	grpc.ClientStream
}

type graphStoreScanClient struct {
	grpc.ClientStream
}

func (x *graphStoreScanClient) Recv() (*kythe_proto.Entry, error) {
	m := new(kythe_proto.Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *graphStoreClient) Write(ctx context.Context, in *kythe_proto.WriteRequest, opts ...grpc.CallOption) (*kythe_proto.WriteReply, error) {
	out := new(kythe_proto.WriteReply)
	err := grpc.Invoke(ctx, "/kythe.proto.GraphStore/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GraphStore service

type GraphStoreServer interface {
	// Read responds with all Entry messages that match the given ReadRequest.
	// The Read operation should be implemented with time complexity proportional
	// to the size of the return set.
	Read(*kythe_proto.ReadRequest, GraphStore_ReadServer) error
	// Scan responds with all Entry messages matching the given ScanRequest.  If a
	// ScanRequest field is empty, any entry value for that field matches and will
	// be returned.  Scan is similar to Read, but with no time complexity
	// restrictions.
	Scan(*kythe_proto.ScanRequest, GraphStore_ScanServer) error
	// Write atomically inserts or updates a collection of entries into the store.
	// Each update is a tuple of the form (kind, target, fact, value).  For each
	// such update, an entry (source, kind, target, fact, value) is written into
	// the store, replacing any existing entry (source, kind, target, fact,
	// value') that may exist.  Note that this operation cannot delete any data
	// from the store; entries are only ever inserted or updated.  Apart from
	// acting atomically, no other constraints are placed on the implementation.
	Write(context.Context, *kythe_proto.WriteRequest) (*kythe_proto.WriteReply, error)
}

func RegisterGraphStoreServer(s *grpc.Server, srv GraphStoreServer) {
	s.RegisterService(&_GraphStore_serviceDesc, srv)
}

func _GraphStore_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(kythe_proto.ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GraphStoreServer).Read(m, &graphStoreReadServer{stream})
}

type GraphStore_ReadServer interface {
	Send(*kythe_proto.Entry) error
	grpc.ServerStream
}

type graphStoreReadServer struct {
	grpc.ServerStream
}

func (x *graphStoreReadServer) Send(m *kythe_proto.Entry) error {
	return x.ServerStream.SendMsg(m)
}

func _GraphStore_Scan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(kythe_proto.ScanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GraphStoreServer).Scan(m, &graphStoreScanServer{stream})
}

type GraphStore_ScanServer interface {
	Send(*kythe_proto.Entry) error
	grpc.ServerStream
}

type graphStoreScanServer struct {
	grpc.ServerStream
}

func (x *graphStoreScanServer) Send(m *kythe_proto.Entry) error {
	return x.ServerStream.SendMsg(m)
}

func _GraphStore_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kythe_proto.WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraphStoreServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kythe.proto.GraphStore/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraphStoreServer).Write(ctx, req.(*kythe_proto.WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GraphStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kythe.proto.GraphStore",
	HandlerType: (*GraphStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _GraphStore_Write_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _GraphStore_Read_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Scan",
			Handler:       _GraphStore_Scan_Handler,
			ServerStreams: true,
		},
	},
}

// Client API for ShardedGraphStore service

type ShardedGraphStoreClient interface {
	// Count returns the number of entries in the given shard.
	Count(ctx context.Context, in *kythe_proto.CountRequest, opts ...grpc.CallOption) (*kythe_proto.CountReply, error)
	// Shard responds with each Entry in the given shard.
	Shard(ctx context.Context, in *kythe_proto.ShardRequest, opts ...grpc.CallOption) (ShardedGraphStore_ShardClient, error)
}

type shardedGraphStoreClient struct {
	cc *grpc.ClientConn
}

func NewShardedGraphStoreClient(cc *grpc.ClientConn) ShardedGraphStoreClient {
	return &shardedGraphStoreClient{cc}
}

func (c *shardedGraphStoreClient) Count(ctx context.Context, in *kythe_proto.CountRequest, opts ...grpc.CallOption) (*kythe_proto.CountReply, error) {
	out := new(kythe_proto.CountReply)
	err := grpc.Invoke(ctx, "/kythe.proto.ShardedGraphStore/Count", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardedGraphStoreClient) Shard(ctx context.Context, in *kythe_proto.ShardRequest, opts ...grpc.CallOption) (ShardedGraphStore_ShardClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ShardedGraphStore_serviceDesc.Streams[0], c.cc, "/kythe.proto.ShardedGraphStore/Shard", opts...)
	if err != nil {
		return nil, err
	}
	x := &shardedGraphStoreShardClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShardedGraphStore_ShardClient interface {
	Recv() (*kythe_proto.Entry, error)
	grpc.ClientStream
}

type shardedGraphStoreShardClient struct {
	grpc.ClientStream
}

func (x *shardedGraphStoreShardClient) Recv() (*kythe_proto.Entry, error) {
	m := new(kythe_proto.Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ShardedGraphStore service

type ShardedGraphStoreServer interface {
	// Count returns the number of entries in the given shard.
	Count(context.Context, *kythe_proto.CountRequest) (*kythe_proto.CountReply, error)
	// Shard responds with each Entry in the given shard.
	Shard(*kythe_proto.ShardRequest, ShardedGraphStore_ShardServer) error
}

func RegisterShardedGraphStoreServer(s *grpc.Server, srv ShardedGraphStoreServer) {
	s.RegisterService(&_ShardedGraphStore_serviceDesc, srv)
}

func _ShardedGraphStore_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kythe_proto.CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardedGraphStoreServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kythe.proto.ShardedGraphStore/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardedGraphStoreServer).Count(ctx, req.(*kythe_proto.CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardedGraphStore_Shard_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(kythe_proto.ShardRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShardedGraphStoreServer).Shard(m, &shardedGraphStoreShardServer{stream})
}

type ShardedGraphStore_ShardServer interface {
	Send(*kythe_proto.Entry) error
	grpc.ServerStream
}

type shardedGraphStoreShardServer struct {
	grpc.ServerStream
}

func (x *shardedGraphStoreShardServer) Send(m *kythe_proto.Entry) error {
	return x.ServerStream.SendMsg(m)
}

var _ShardedGraphStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kythe.proto.ShardedGraphStore",
	HandlerType: (*ShardedGraphStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Count",
			Handler:    _ShardedGraphStore_Count_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Shard",
			Handler:       _ShardedGraphStore_Shard_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptorStorageService = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0xae, 0x2c, 0xc9,
	0x48, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0x8d,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0x8b, 0x0a, 0x71, 0x83, 0x95, 0x40, 0x38,
	0x52, 0x92, 0x58, 0xd4, 0x43, 0xa4, 0x8c, 0xf6, 0x33, 0x72, 0x71, 0xb9, 0x17, 0x25, 0x16, 0x64,
	0x04, 0x97, 0xe4, 0x17, 0xa5, 0x0a, 0x59, 0x70, 0xb1, 0x04, 0xa5, 0x26, 0xa6, 0x08, 0x49, 0xe8,
	0x21, 0xe9, 0xd7, 0x03, 0x09, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0x09, 0xa1, 0xc8,
	0xb8, 0xe6, 0x95, 0x14, 0x55, 0x2a, 0x31, 0x18, 0x30, 0x82, 0x74, 0x06, 0x27, 0x27, 0xe6, 0xa1,
	0xe9, 0x04, 0x09, 0x11, 0xd2, 0x69, 0xcb, 0xc5, 0x1a, 0x5e, 0x94, 0x59, 0x92, 0x2a, 0x24, 0x89,
	0xa2, 0x00, 0x2c, 0x06, 0xd3, 0x2b, 0x8e, 0x4d, 0xaa, 0x20, 0xa7, 0x52, 0x89, 0xc1, 0xa8, 0x8f,
	0x91, 0x4b, 0x30, 0x38, 0x23, 0xb1, 0x28, 0x25, 0x35, 0x05, 0xc9, 0x23, 0xb6, 0x5c, 0xac, 0xce,
	0xf9, 0xa5, 0x79, 0x25, 0x68, 0x86, 0x82, 0xc5, 0xb0, 0x1b, 0x0a, 0x95, 0x02, 0x1b, 0x2a, 0x64,
	0xc5, 0xc5, 0x0a, 0x36, 0x13, 0x4d, 0x3b, 0x58, 0x8c, 0x80, 0x7f, 0x9c, 0x0c, 0x4f, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xb8, 0xe4,
	0x93, 0xf3, 0x73, 0xf5, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x52, 0x52, 0xcb, 0x4a, 0xf2,
	0xf3, 0x73, 0x8a, 0x91, 0x35, 0x27, 0xb1, 0x81, 0x29, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x58, 0x32, 0x4e, 0x26, 0xd9, 0x01, 0x00, 0x00,
}