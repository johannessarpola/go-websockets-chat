// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

	cp := NewChatProducer(client, "name", "tests")
	cc := NewChatConsumer(client, "name", "tests")

	go cp.Run()
	go cc.Run()

	fmt.Println("Published message")

	flag.Parse()
	// hub := NewHub()

	u := NewUser("name")
	m := NewMessage(u, "mesage")

	cp.in <- *m

	//	go hub.run() // TODO Reconfigure to use Pulsar
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {

		//		ServeWebsocket(hub, w, r)
	})

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
