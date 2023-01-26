package app

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
type ChatProducer struct {
	// Inbound messages from the clients.
	Channel        chan models.Message
	pulsarProducer pulsar.Producer
	name           string
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
		Channel:        make(chan models.Message),
		pulsarProducer: producer,
		name:           name,
	}
}

func (cp *ChatProducer) Close() {
	cp.pulsarProducer.Close()
}

func (cp *ChatProducer) Run() {

	for message := range cp.Channel {

		fmt.Println("internal pipe")
		fmt.Println(message.Message)

		pd, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = cp.pulsarProducer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: pd,
		})
		println("sent to pulsars")
		if err != nil {
			println("err err")
			fmt.Println(err)
		}

	}
}
