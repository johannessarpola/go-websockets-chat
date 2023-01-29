// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/johannessarpola/go-websockets-chat/messaging"
	"github.com/johannessarpola/go-websockets-chat/models"
	"github.com/johannessarpola/go-websockets-chat/server"
	"google.golang.org/grpc/reflection"
)

func App() {
	flag.Parse()

	gw := messaging.NewPulsarGateway("pulsar://localhost:6650")
	go gw.Run()
	go genSomething(gw)
	go printMsgs(gw)

	srv := server.NewGRPCServer()
	port := ":8080"

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

func printMsgs(gw messaging.Gateway) {
	count := 0

	for {
		msgs, _ := gw.Poll()
		fmt.Println("Received batch of messages")
		for _, m := range msgs {
			count++
			fmt.Printf("Received message %s\n", m.Message)
		}

		fmt.Printf("Received %d messages\n", count)
	}
}

func genSomething(gw messaging.Gateway) {
	u := models.NewUser("heartbeat")
	for {
		msgBody := fmt.Sprintf("message genrated at %s", time.Now())
		m := models.NewMessage(u, msgBody)
		gw.Send(*m)
		time.Sleep(1 * time.Second)
	}
}
