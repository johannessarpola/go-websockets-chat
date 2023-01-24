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
	pulsarConsumer pulsar.Consumer
	Channel        chan Message
	name           string
}

func NewChatConsumer(client pulsar.Client, name string, topic string, dlq string) *ChatConsumer {
	//dlqOptions := pulsar.DLQPolicy{
	//	DeadLetterTopic: dlq,
	//	MaxDeliveries:   1,
	// TODO RetryLetterTopic
	//}

	options := pulsar.ConsumerOptions{
		// TODO Schema
		Topic:            topic,
		SubscriptionName: name,
		Type:             pulsar.Shared,
		//DLQ:              &dlqOptions,
	}

	consumer, err := client.Subscribe(options)

	if err != nil {
		log.Fatalf("Could not start consumer: %v", err)
	}

	return &ChatConsumer{
		pulsarConsumer: consumer,
		Channel:        make(chan Message),
		name:           name,
	}
}

func (cc *ChatConsumer) Close() {
	cc.pulsarConsumer.Close()
}

func (cc *ChatConsumer) Run() {

	for record := range cc.pulsarConsumer.Chan() {
		msg := record.Message
		cc.pulsarConsumer.Ack(msg)

		fmt.Println("internal consumer")
		jsonMsg := Message{}

		err := json.Unmarshal(msg.Payload(), &jsonMsg)
		if err != nil {
			println("err err")
			log.Println(err)
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), jsonMsg.Message)

		//cc.Channel <- jsonMsg // TODO Wire up to front-end
	}
}
