package app

import (
	//"context"
	//"log"
	//"time"

	api "github.com/johannessarpola/go-websockets-chat/api/v1"
	"google.golang.org/grpc"
	//"google.golang.org/grpc"
	//codes "google.golang.org/grpc/codes"
	//status "google.golang.org/grpc/status"
)

type grpcServer struct {
	api.UnimplementedChatServer
	// TODO Pulsar producer
	// TODO Pulsar consumer
}

func NewGRPCServer() *grpc.Server {

	//var err error
	gsrv := grpc.NewServer()
	srv := grpcServer{
		// TODO Pulsar producer
		// TODO Pulsar consumer
	}
	api.RegisterChatServer(gsrv, &srv)
	return gsrv
}
