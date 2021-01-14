package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/google/uuid"
	"syreclabs.com/go/faker"
)

const (
	pulsarBrokerURL  = "pulsar://localhost:6650"
	pulsarTopic      = "my_topic"
	pulsarOpTimeout  = 30 * time.Second
	pulsarConTimeout = 30 * time.Second
	firingDelay      = 3 * time.Second
)

func main() {
	// --- Signal handler for a more graceful shutdown ---
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// --- Setup ---

	// Set up client.
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               pulsarBrokerURL,
		OperationTimeout:  pulsarOpTimeout,
		ConnectionTimeout: pulsarConTimeout,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	// Set up subscriber.
	cMessages := make(chan pulsar.ConsumerMessage, 10)
	sub, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            pulsarTopic,
		SubscriptionName: fmt.Sprintf("%s-subscriber", pulsarTopic),
		Type:             pulsar.Shared,
		MessageChannel:   cMessages,
	})
	if err != nil {
		panic(err)
	}

	defer sub.Close()

	// Set up a publisher.
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: pulsarTopic,
	})
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	// --- Asynchronous processing of outgoing messages ---

	// Process incoming messages received via go channel.
	go consumeMessages(sub, cMessages)

	// Process outgoing messages.
	go publishMessages(producer)

	sig := <-sigChan
	log.Printf("received signal %d, terminating...\n", sig)
	os.Exit(0)
}

func consumeMessages(sub pulsar.Consumer, cMessages <-chan pulsar.ConsumerMessage) {
	for cmsg := range cMessages {
		msg := cmsg.Message
		log.Printf("received message: %s, payload: %s", msg.ID(), string(msg.Payload()))

		// Acknowledge that the message is received otherwise it will be resent
		// repeatedly.
		sub.Ack(msg)
	}
}

func publishMessages(producer pulsar.Producer) {
	for {
		id := uuid.New().String()
		text := fmt.Sprintf("Hello %s", faker.Name().Name())
		msg := pulsar.ProducerMessage{
			Key:     id,
			Payload: []byte(text),
		}
		log.Printf("sent message: %s, payload: %s", id, text)
		_, err := producer.Send(context.Background(), &msg)
		if err != nil {
			panic(err)
		}

		time.Sleep(firingDelay)
	}
}
