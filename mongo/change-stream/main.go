package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"syreclabs.com/go/faker"
	"syscall"
	"time"
)

const (
	firingDelay        = 3 * time.Second
)

func getLocalDSN() string {
	return "mongodb://nobody:secrets@localhost:27017,localhost:27018,localhost:27019/test?replicaSet=rs0&connect=direct"
}

func main() {
	// --- Signal handler for a more graceful shutdown ---
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// --- Setup ---

	// Set up connection to Mongo.
	dsn := getLocalDSN()
	client, err := mongo.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("can't connect to mongo: %v", err)
	}

	mongoCtx, cancel := context.WithCancel(context.Background())
	defer func() {
		client.Disconnect(mongoCtx)
		cancel()
	}()
	if err := client.Connect(mongoCtx); err != nil {
		log.Fatalf("can't connect to mongo: %v", err)
	}

	// --- Run publisher and subscriber concurrently ---

	// Process incoming messages received via go channel.
	go processFromMongo(mongoCtx, client)

	// Process outgoing messages.
	go publishToMongo(mongoCtx, client)

	sig := <-sigChan
	log.Printf("received signal %d, terminating...\n", sig)
	mongoCtx.Done()
	os.Exit(0)
}

func publishToMongo(ctx context.Context, client *mongo.Client) {
	db := client.Database("test")
	messages := db.Collection("messages")

	for {
		msg := fmt.Sprintf("Hello %s", faker.Name().Name())

		result, err := messages.InsertOne(ctx, bson.D{
			{Key: "msg", Value: msg},
		})
		if err != nil {
			log.Fatalf("can't insert %s to mongo: %v", msg, err)
		}

		log.Printf("successfully inserted %s to mongo with id: %s", msg, result.InsertedID)

		time.Sleep(firingDelay)
	}
}


func processFromMongo(ctx context.Context, client *mongo.Client) {
	db := client.Database("test")
	messages := db.Collection("messages")

	cursor, err := messages.Watch(ctx, mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		log.Fatalf("can't start change stream: %v", err)
	}
	defer func() {
		cursor.Close(ctx)
	}()

	for {
		if ok := cursor.Next(ctx); !ok {
			continue
		}
		next := cursor.Current
		log.Printf("current: %v", next)
	}
}
