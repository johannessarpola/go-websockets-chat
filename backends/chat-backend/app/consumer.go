package app

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type ChatConsumer struct {
	// Inbound messages from the clients.
	in       chan pulsar.ConsumerMessage
	consumer pulsar.Consumer
	out      chan Message
	name     string
}

func NewChatConsumer(client pulsar.Client, name string, topic string, dlq string) *ChatConsumer {

	channel := make(chan pulsar.ConsumerMessage)

	dlqOptions := pulsar.DLQPolicy{
		DeadLetterTopic: dlq,
		// TODO RetryLetterTopic
	}

	options := pulsar.ConsumerOptions{
		// TODO Schema
		Topic:            topic,
		SubscriptionName: name,
		Type:             pulsar.Shared,
		DLQ:              &dlqOptions,
	}

	options.MessageChannel = channel

	consumer, err := client.Subscribe(options)

	if err != nil {
		log.Fatalf("Could not start consumer: %v", err)
	}

	return &ChatConsumer{
		in:       channel,
		consumer: consumer,
		out:      make(chan Message),
		name:     name,
	}
}

func (cc *ChatConsumer) Run() {

	defer cc.consumer.Close()

	for cm := range cc.in {
		msg := cm.Message
		jsonMsg := Message{}

		err := json.Unmarshal(msg.Payload(), &jsonMsg)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), jsonMsg.Message)

		cc.out <- jsonMsg
		cc.consumer.Ack(msg)
	}
}
