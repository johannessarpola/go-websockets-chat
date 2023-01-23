// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"flag"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

var addr = flag.String("addr", ":8080", "http service address")

func App() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	cp := NewChatProducer(client, "producer", "tests")
	cc := NewChatConsumer(client, "consumer", "tests")

	go cp.Run()
	go cc.Run()

	flag.Parse()
	// hub := NewHub()

	u := NewUser("name")
	m := NewMessage(u, "message")

	m2 := NewMessage(u, "message2")

	cp.in <- *m
	cp.in <- *m2

	//	go hub.run() // TODO Reconfigure to use Pulsar
	//http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {

	//		ServeWebsocket(hub, w, r)
	//})

	//err = http.ListenAndServe(*addr, nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
