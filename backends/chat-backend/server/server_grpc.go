package server

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
	//toBackend   chan models.Message
	//fromBackend chan models.Message
}

func NewGRPCServer( /* toBackend chan models.Message, fromBackend chan models.Message */ ) *grpc.Server {

	//var err error
	gsrv := grpc.NewServer()
	srv := grpcServer{
		// toBackend:   toBackend,
		// fromBackend: fromBackend,
	}
	api.RegisterChatServer(gsrv, &srv)
	return gsrv
}
