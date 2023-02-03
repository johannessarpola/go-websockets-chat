// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/v1/chat.proto

package api_v1

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	Register(ctx context.Context, in *UserRegister, opts ...grpc.CallOption) (*UserRegisterReply, error)
	ListUsers(ctx context.Context, in *Null, opts ...grpc.CallOption) (*UserListReply, error)
	Message(ctx context.Context, in *NewMessage, opts ...grpc.CallOption) (*Null, error)
	Poll(ctx context.Context, in *Null, opts ...grpc.CallOption) (*PollReply, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Register(ctx context.Context, in *UserRegister, opts ...grpc.CallOption) (*UserRegisterReply, error) {
	out := new(UserRegisterReply)
	err := c.cc.Invoke(ctx, "/api.v1.Chat/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) ListUsers(ctx context.Context, in *Null, opts ...grpc.CallOption) (*UserListReply, error) {
	out := new(UserListReply)
	err := c.cc.Invoke(ctx, "/api.v1.Chat/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Message(ctx context.Context, in *NewMessage, opts ...grpc.CallOption) (*Null, error) {
	out := new(Null)
	err := c.cc.Invoke(ctx, "/api.v1.Chat/Message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Poll(ctx context.Context, in *Null, opts ...grpc.CallOption) (*PollReply, error) {
	out := new(PollReply)
	err := c.cc.Invoke(ctx, "/api.v1.Chat/Poll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	Register(context.Context, *UserRegister) (*UserRegisterReply, error)
	ListUsers(context.Context, *Null) (*UserListReply, error)
	Message(context.Context, *NewMessage) (*Null, error)
	Poll(context.Context, *Null) (*PollReply, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) Register(context.Context, *UserRegister) (*UserRegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedChatServer) ListUsers(context.Context, *Null) (*UserListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedChatServer) Message(context.Context, *NewMessage) (*Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Message not implemented")
}
func (UnimplementedChatServer) Poll(context.Context, *Null) (*PollReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Poll not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Chat/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Register(ctx, req.(*UserRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Chat/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).ListUsers(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Chat/Message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Message(ctx, req.(*NewMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Poll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Poll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.Chat/Poll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Poll(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Chat_Register_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Chat_ListUsers_Handler,
		},
		{
			MethodName: "Message",
			Handler:    _Chat_Message_Handler,
		},
		{
			MethodName: "Poll",
			Handler:    _Chat_Poll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/chat.proto",
}