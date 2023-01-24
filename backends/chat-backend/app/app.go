// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"flag"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

func App() {

	gw := NewPulsarGateway("pulsar://localhost:6650")
	go gw.Run()

	flag.Parse()
	// hub := NewHub()

	u := NewUser("name")
	m := NewMessage(u, "message")
	m2 := NewMessage(u, "message2")
	m3 := NewMessage(u, "message3")

	gw.Producer.Channel <- *m
	time.Sleep(5 * time.Second)
	gw.Producer.Channel <- *m2
	time.Sleep(5 * time.Second)
	gw.Producer.Channel <- *m3
	time.Sleep(5 * time.Second)

	for {
		time.Sleep(time.Second)
	}

	//	go hub.run() // TODO Reconfigure to use Pulsar
	//http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {

	//		ServeWebsocket(hub, w, r)
	//})

	//err = http.ListenAndServe(*addr, nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
