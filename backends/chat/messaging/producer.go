package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/johannessarpola/go-websockets-chat/models"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type chatProducer struct {
	// Inbound messages from the clients.
	Channel        chan models.Message
	pulsarProducer pulsar.Producer
	name           string
}

func NewChatProducer(client pulsar.Client, name string, topic string) *chatProducer {

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
		Name:  name,
	})

	if err != nil {
		log.Fatalf("Could not start producer: %v", err)
	}

	return &chatProducer{
		Channel:        make(chan models.Message),
		pulsarProducer: producer,
		name:           name,
	}
}

func (cp *chatProducer) Close() {
	cp.pulsarProducer.Close()
}

func (cp *chatProducer) Run() {

	for message := range cp.Channel {
		pd, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err) // TODO
			return
		}
		_, err = cp.pulsarProducer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: pd,
		})
		println("\nsent to pulsars")
		if err != nil {
			println("err err")
			fmt.Println(err)
		}

	}
}
