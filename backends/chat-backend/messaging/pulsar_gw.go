package messaging

import (
	"errors"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/johannessarpola/go-websockets-chat/models"
)

type PulsarGateway struct {
	client   pulsar.Client
	producer *chatProducer
	consumer *chatConsumer
}

func (gw *PulsarGateway) Send(message models.Message) error {
	gw.producer.Channel <- message
	return nil // TODO
}

func (gw *PulsarGateway) Poll() ([]models.Message, error) {
	// TODO Implement
	return nil, errors.New("Not implemented, yet")
}

// Event loop with cleanup
func (gw *PulsarGateway) Run() {
	defer gw.client.Close()
	defer gw.producer.Close()
	defer gw.consumer.Close()

	go gw.consumer.Run()
	gw.producer.Run()

}

func NewPulsarGateway(addr string) Gateway {
	// TODO Configure
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               addr,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Fatalf("Could not create pulsar client: %v", err)
	}

	// TODO Configure
	cp := NewChatProducer(client, "producer", "tests")
	cc := NewChatConsumer(client, "consumer", "tests", "dlq")

	var gw Gateway = &PulsarGateway{
		client:   client,
		producer: cp,
		consumer: cc,
	}
	return gw

}
