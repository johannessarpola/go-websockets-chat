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
type ChatProducer struct {
	// Inbound messages from the clients.
	in   chan Message
	out  pulsar.Producer
	name string
}

func NewChatProducer(client pulsar.Client, name string, topic string) *ChatProducer {

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
		Name:  name,
	})

	if err != nil {
		log.Fatalf("Could not start producer: %v", err)
	}

	return &ChatProducer{
		in:   make(chan Message),
		out:  producer,
		name: name,
	}
}

func (cp *ChatProducer) Run() {

	defer cp.out.Close()

	for nessage := range cp.in {

		fmt.Println("internal pipe")
		fmt.Println(nessage.Message)

		pd, err := json.Marshal(nessage)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = cp.out.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: pd,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
