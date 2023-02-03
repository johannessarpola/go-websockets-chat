package server

import (
	//"context"
	//"log"
	//"time"

	"context"

	api "github.com/johannessarpola/go-websockets-chat/api/v1"
	"github.com/johannessarpola/go-websockets-chat/messaging"
	"github.com/johannessarpola/go-websockets-chat/models"
	"github.com/johannessarpola/go-websockets-chat/utils"
	"google.golang.org/grpc"
	//"google.golang.org/grpc"
	//codes "google.golang.org/grpc/codes"
	//status "google.golang.org/grpc/status"
)

type grpcServer struct {
	api.UnimplementedChatServer
	gw messaging.Gateway
}

func (s *grpcServer) Register(context.Context, *api.UserRegister) (*api.UserRegisterReply, error) {
	return nil, nil
}

func (s *grpcServer) ListUsers(context.Context, *api.Null) (*api.UserListReply, error) {
	return nil, nil
}

func (s *grpcServer) Message(context.Context, *api.NewMessage) (*api.Null, error) {
	return nil, nil
}

func transform(m models.Message) *api.Message {
	u := api.User{
		Id:   m.User.Id,
		Name: m.User.Name,
	}

	return &api.Message{
		User:    &u,
		Time:    nil, // TODO Transform
		Message: m.Message,
	}
}

func (s *grpcServer) Poll(context.Context, *api.Null) (*api.PollReply, error) {
	arr, _ := s.gw.Poll() // TODO Error handling
	r := utils.Map(arr, transform)

	return &api.PollReply{
		Messages: r,
	}, nil

}

//rpc Register (UserRegister) returns (UserRegisterReply) {}
//rpc ListUsers (Null) returns (UserListReply) {}
//rpc Message (NewMessage) returns (Null) {}
//rpc Poll (Null) returns (PollReply) {}

func NewGRPCServer(gw messaging.Gateway) *grpc.Server {

	//var err error
	gsrv := grpc.NewServer()
	srv := grpcServer{gw: gw}
	api.RegisterChatServer(gsrv, &srv)
	return gsrv
}
