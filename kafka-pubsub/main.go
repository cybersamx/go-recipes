// Based on  https://watermill.io/docs/getting-started/

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"syreclabs.com/go/faker"
)

const (
	kafkaBrokerURL     = "localhost:9092"
	kafkaConsumerGroup = "consumer_group"
	kafkaTopic         = "my_topic"
	firingDelay        = 3 * time.Second
)

var (
	logger = watermill.NewStdLogger(false, false)
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

	// Set up subscriber.
	saramaConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	subConfig := kafka.SubscriberConfig{
		Brokers:               []string{kafkaBrokerURL},
		Unmarshaler:           kafka.DefaultMarshaler{},
		OverwriteSaramaConfig: saramaConfig,
		ConsumerGroup:         kafkaConsumerGroup,
	}

	sub, err := kafka.NewSubscriber(subConfig, logger)
	if err != nil {
		panic(err)
	}

	// Set up publisher.
	pubConfig := kafka.PublisherConfig{
		Brokers:   []string{"localhost:9092"},
		Marshaler: kafka.DefaultMarshaler{},
	}

	pub, err := kafka.NewPublisher(pubConfig, logger)
	if err != nil {
		panic(err)
	}

	// --- Asynchronous processing of incoming and outgoing messages ---

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	messages, err := sub.Subscribe(ctx, kafkaTopic)
	if err != nil {
		panic(err)
	}

	// Process incoming messages received via a go channel.
	go consumeMessages(messages)

	// Process outgoing messages.
	go publishMessages(pub)

	sig := <-sigChan
	log.Printf("received signal %d, terminating...\n", sig)
	ctx.Done()
	os.Exit(0)
}

func consumeMessages(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// Acknowledge that the message is received otherwise it will be resent
		// repeatedly.
		msg.Ack()
	}
}

func publishMessages(pub message.Publisher) {
	for {
		uuid := watermill.NewUUID()
		msg := message.NewMessage(uuid, []byte(fmt.Sprintf("Hello %s", faker.Name().Name())))
		log.Printf("sent message: %s, payload: %s", uuid, string(msg.Payload))

		if err := pub.Publish(kafkaTopic, msg); err != nil {
			panic(err)
		}

		time.Sleep(firingDelay)
	}
}
