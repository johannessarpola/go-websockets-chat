package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/johannessarpola/go-websockets-chat/models"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type chatConsumer struct {
	// Inbound messages from the clients.
	pulsarConsumer pulsar.Consumer
	name           string
}

func NewChatConsumer(client pulsar.Client, name string, topic string, dlq string) *chatConsumer {
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

	return &chatConsumer{
		pulsarConsumer: consumer,
		name:           name,
	}
}

func (cc *chatConsumer) Close() {
	cc.pulsarConsumer.Close()
}

func transformMessage(pmsg pulsar.Message) models.Message {
	jsonMessage := models.Message{}
	err := json.Unmarshal(pmsg.Payload(), &jsonMessage)
	if err != nil {
		log.Fatal(err)
	}
	return jsonMessage
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func nonNil(s models.Message) bool {
	return s != models.Message{}
}

func (cc *chatConsumer) Poll() ([]models.Message, error) {

	size := 10
	arr := make([]models.Message, size)

	// channel := cc.pulsarConsumer.Chan()

	// select {
	// case <-time.After(10 * time.Second):
	// case pm := <-channel:
	// 	fmt.Printf("Message ID of %s\n", pm.ID())
	// 	arr = append(arr, transformMessage(pm))
	// 	cc.pulsarConsumer.Ack(pm)
	// }

	start := time.Now()

	for {

		// Rudimentary, buffer for max 1 second messages from backend
		if !start.Add(1 * time.Second).After(time.Now()) {
			break
		}
		if size == 0 {
			break
		}
		pm, _ := cc.pulsarConsumer.Receive(context.Background())
		arr = append(arr, transformMessage(pm))
		cc.pulsarConsumer.Ack(pm)
		size--
	}

	fmt.Println("Acked last")
	return filter(arr, nonNil), nil
}

// func (cc *chatConsumer) Run() {

// 	for record := range cc.pulsarConsumer.Chan() {
// 		msg := record.Message
// 		cc.pulsarConsumer.Ack(msg)

// 		fmt.Println("internal consumer")
// 		jsonMsg := models.Message{}

// 		err := json.Unmarshal(msg.Payload(), &jsonMsg)
// 		if err != nil {
// 			println("err err")
// 			log.Println(err)
// 		}

// 		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
// 			msg.ID(), jsonMsg.Message)

// 		//cc.Channel <- jsonMsg // TODO Wire up to front-end
// 	}
// }
