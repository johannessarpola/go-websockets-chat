package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type ChatConsumer struct {
	// Inbound messages from the clients.
	in   pulsar.Consumer
	out  chan Message
	name string
}

func NewChatConsumer(client pulsar.Client, name string, topic string) *ChatConsumer {

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: name,
	})

	if err != nil {
		log.Fatalf("Could not start consumer: %v", err)
	}

	return &ChatConsumer{
		in:   consumer,
		out:  make(chan Message),
		name: name,
	}
}

func (cc *ChatConsumer) Run() {
	for {
		evt, err := cc.in.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		msg := Message{}

		err = json.Unmarshal(evt.Payload(), &msg)
		if err != nil {
			// TODO Send to DLQ/or ignore
			log.Fatal(err)
		}
		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			evt.ID(), string(evt.Payload()))

		cc.out <- msg
	}
}
