package app

import (
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarGateway struct {
	Client   pulsar.Client
	Producer *ChatProducer
	Consumer *ChatConsumer
}

// Event loop with cleanup
func (gw *PulsarGateway) Run() {
	defer gw.Client.Close()
	defer gw.Producer.Close()
	defer gw.Consumer.Close()

	go gw.Consumer.Run()
	go gw.Producer.Run()

	for {
		time.Sleep(5 * time.Second)
	}
}

func NewPulsarGateway(addr string) *PulsarGateway {
	// TODO Configure ?
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

	return &PulsarGateway{
		Client:   client,
		Producer: cp,
		Consumer: cc,
	}

}
