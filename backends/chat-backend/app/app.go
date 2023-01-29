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

func App() {
	flag.Parse()
	//mux := http.NewServeMux()
	//reflection.Register(srv)

	//mux.HandleFunc("/", srv.ServeHTTP)

	gw := NewPulsarGateway("pulsar://localhost:6650")
	go gw.Run()
	go genSomething(gw)

	srv := server.NewGRPCServer()
	port := ":8080"
	//http.ListenAndServe(":8082", mux)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(srv) // TODO If debug use this
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer lis.Close()
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
