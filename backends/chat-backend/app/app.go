// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

// Skills: Software Development · OpenShift · Groovy · Grails · Docker · Java · JavaScript
import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

var addr = flag.String("addr", ":8080", "http service address")

func startProducer(client pulsar.Client, topic string) {

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	if err != nil {
		log.Fatalf("Could not start producer: %v", err)
	}

	defer producer.Close()

	for {
		time.Sleep(5 * time.Second) // TODO Remove

		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte("hello world"),
		})

		if err != nil {
			fmt.Println("Failed to publish message", err)
		}
	}
}

func startConsumer(client pulsar.Client, subscription string, topic string) {

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subscription,
		Type:             pulsar.Shared,
	})

	if err != nil {
		log.Fatalf("Could not subcscribe consumer: %v", err)
	}

	defer consumer.Close()

	msg, err := consumer.Receive(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
		msg.ID(), string(msg.Payload()))

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))
	}
}

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

	go startProducer(client, "test-topic")
	go startConsumer(client, "sub", "test-topic")

	fmt.Println("Published message")

	flag.Parse()
	hub := newHub()

	go hub.run() // TODO Reconfigure to use Pulsar
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(hub, w, r)
	})

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
