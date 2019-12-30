package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
	"log"
	"time"
)

const timeout = 2 * time.Second
const retries = 5

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Some arbitrary URI string
	uri := "mongodb://user:password@localhost:27017/database"
	client, _ := mongo.NewClient(options.Client().ApplyURI(uri))
	defer client.Disconnect(ctx)

	for i := 0; i < retries; i++ {
		log.Printf("attempt %d: connecting to mongo %s", i + 1, uri)
		cctx, _ := context.WithTimeout(ctx, timeout)
		err := client.Connect(cctx)
		if err != nil && err != topology.ErrTopologyConnected {
			log.Printf("error: %v", err)
		}

		rctx, _ := context.WithTimeout(ctx, timeout)
		if err := client.Ping(rctx, readpref.Primary()); err != nil {
			// Ping mongo failed, retry unless we exhausted our retries
			if i == retries-1 {
				panic("exhausted our retries")
			}

			continue // Retry
		} else {
			log.Printf("connected to %s successfully\n", uri)
			break
		}
	}
}
