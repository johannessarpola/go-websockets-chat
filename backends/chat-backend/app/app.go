// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/johannessarpola/go-websockets-chat/models"
	"github.com/johannessarpola/go-websockets-chat/server"
	"google.golang.org/grpc/reflection"
)

var addr = flag.String("addr", ":8080", "http service address")

func App() {

	gw := NewPulsarGateway("pulsar://localhost:6650")
	go gw.Run()
	go serveHttp(":8082")
	go genSomething(gw)

	flag.Parse()

	for {
		time.Sleep(time.Second)
	}

}

func genSomething(gw *PulsarGateway) {
	u := models.NewUser("name")
	m := models.NewMessage(u, "message")
	m2 := models.NewMessage(u, "message2")
	m3 := models.NewMessage(u, "message3")

	gw.Producer.Channel <- *m
	time.Sleep(5 * time.Second)
	gw.Producer.Channel <- *m2
	time.Sleep(5 * time.Second)
	gw.Producer.Channel <- *m3
	time.Sleep(5 * time.Second)
}

func serveHttp(port string) {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer lis.Close()

	srv := server.NewGRPCServer()

	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
