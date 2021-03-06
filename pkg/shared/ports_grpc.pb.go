// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shared

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

// StreamerClient is the client API for Streamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamerClient interface {
	Send(ctx context.Context, opts ...grpc.CallOption) (Streamer_SendClient, error)
}

type streamerClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamerClient(cc grpc.ClientConnInterface) StreamerClient {
	return &streamerClient{cc}
}

func (c *streamerClient) Send(ctx context.Context, opts ...grpc.CallOption) (Streamer_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &Streamer_ServiceDesc.Streams[0], "/shared.Streamer/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerSendClient{stream}
	return x, nil
}

type Streamer_SendClient interface {
	Send(*PortBlueprint) error
	CloseAndRecv() (*Summary, error)
	grpc.ClientStream
}

type streamerSendClient struct {
	grpc.ClientStream
}

func (x *streamerSendClient) Send(m *PortBlueprint) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamerSendClient) CloseAndRecv() (*Summary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Summary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamerServer is the server API for Streamer service.
// All implementations must embed UnimplementedStreamerServer
// for forward compatibility
type StreamerServer interface {
	Send(Streamer_SendServer) error
	mustEmbedUnimplementedStreamerServer()
}

// UnimplementedStreamerServer must be embedded to have forward compatible implementations.
type UnimplementedStreamerServer struct {
}

func (UnimplementedStreamerServer) Send(Streamer_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedStreamerServer) mustEmbedUnimplementedStreamerServer() {}

// UnsafeStreamerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamerServer will
// result in compilation errors.
type UnsafeStreamerServer interface {
	mustEmbedUnimplementedStreamerServer()
}

func RegisterStreamerServer(s grpc.ServiceRegistrar, srv StreamerServer) {
	s.RegisterService(&Streamer_ServiceDesc, srv)
}

func _Streamer_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamerServer).Send(&streamerSendServer{stream})
}

type Streamer_SendServer interface {
	SendAndClose(*Summary) error
	Recv() (*PortBlueprint, error)
	grpc.ServerStream
}

type streamerSendServer struct {
	grpc.ServerStream
}

func (x *streamerSendServer) SendAndClose(m *Summary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamerSendServer) Recv() (*PortBlueprint, error) {
	m := new(PortBlueprint)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Streamer_ServiceDesc is the grpc.ServiceDesc for Streamer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Streamer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shared.Streamer",
	HandlerType: (*StreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _Streamer_Send_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/ports.proto",
}
