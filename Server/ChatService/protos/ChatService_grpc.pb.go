// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: protos/ChatService.proto

package ChatService_proto

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
	ChatService_CreateChat_FullMethodName   = "/ChatService.ChatService/CreateChat"
	ChatService_SendMessage_FullMethodName  = "/ChatService.ChatService/SendMessage"
	ChatService_AddMember_FullMethodName    = "/ChatService.ChatService/AddMember"
	ChatService_RemoveMember_FullMethodName = "/ChatService.ChatService/RemoveMember"
	ChatService_ListMembers_FullMethodName  = "/ChatService.ChatService/ListMembers"
	ChatService_ListMsgs_FullMethodName     = "/ChatService.ChatService/ListMsgs"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreateChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Void, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Msg, Msg], error)
	AddMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Msg, error)
	RemoveMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Msg, error)
	ListMembers(ctx context.Context, in *Chat, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Member], error)
	ListMsgs(ctx context.Context, in *Chat, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Msg], error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ChatService_CreateChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Msg, Msg], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_SendMessage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Msg, Msg]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_SendMessageClient = grpc.BidiStreamingClient[Msg, Msg]

func (c *chatServiceClient) AddMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Msg)
	err := c.cc.Invoke(ctx, ChatService_AddMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RemoveMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Msg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Msg)
	err := c.cc.Invoke(ctx, ChatService_RemoveMember_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ListMembers(ctx context.Context, in *Chat, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Member], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], ChatService_ListMembers_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Chat, Member]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ListMembersClient = grpc.ServerStreamingClient[Member]

func (c *chatServiceClient) ListMsgs(ctx context.Context, in *Chat, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Msg], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[2], ChatService_ListMsgs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Chat, Msg]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ListMsgsClient = grpc.ServerStreamingClient[Msg]

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility.
type ChatServiceServer interface {
	CreateChat(context.Context, *Chat) (*Void, error)
	SendMessage(grpc.BidiStreamingServer[Msg, Msg]) error
	AddMember(context.Context, *Member) (*Msg, error)
	RemoveMember(context.Context, *Member) (*Msg, error)
	ListMembers(*Chat, grpc.ServerStreamingServer[Member]) error
	ListMsgs(*Chat, grpc.ServerStreamingServer[Msg]) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) CreateChat(context.Context, *Chat) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(grpc.BidiStreamingServer[Msg, Msg]) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) AddMember(context.Context, *Member) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMember not implemented")
}
func (UnimplementedChatServiceServer) RemoveMember(context.Context, *Member) (*Msg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMember not implemented")
}
func (UnimplementedChatServiceServer) ListMembers(*Chat, grpc.ServerStreamingServer[Member]) error {
	return status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}
func (UnimplementedChatServiceServer) ListMsgs(*Chat, grpc.ServerStreamingServer[Msg]) error {
	return status.Errorf(codes.Unimplemented, "method ListMsgs not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}
func (UnimplementedChatServiceServer) testEmbeddedByValue()                     {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*Chat))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMessage(&grpc.GenericServerStream[Msg, Msg]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_SendMessageServer = grpc.BidiStreamingServer[Msg, Msg]

func _ChatService_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_AddMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddMember(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RemoveMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RemoveMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_RemoveMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RemoveMember(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ListMembers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Chat)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ListMembers(m, &grpc.GenericServerStream[Chat, Member]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ListMembersServer = grpc.ServerStreamingServer[Member]

func _ChatService_ListMsgs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Chat)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ListMsgs(m, &grpc.GenericServerStream[Chat, Msg]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_ListMsgsServer = grpc.ServerStreamingServer[Msg]

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatService.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "AddMember",
			Handler:    _ChatService_AddMember_Handler,
		},
		{
			MethodName: "RemoveMember",
			Handler:    _ChatService_RemoveMember_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _ChatService_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListMembers",
			Handler:       _ChatService_ListMembers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListMsgs",
			Handler:       _ChatService_ListMsgs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/ChatService.proto",
}
