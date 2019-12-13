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

const Timeout = 5 * time.Second

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	ctx := context.Background()

	// Some arbitrary URI string
	uri := "mongodb://user:password@localhost:27017/database"
	client, _ := mongo.NewClient(options.Client().ApplyURI(uri))
	defer client.Disconnect(ctx)

	limit := 10
	for i := 0; i < limit; i++ {
		log.Printf("attempt %d: connecting to mongo %s", i + 1, uri)
		cctx, _ := context.WithTimeout(ctx, Timeout)
		err := client.Connect(cctx)
		if err != nil && err != topology.ErrTopologyConnected {
			log.Printf("error: %v", err)
		}

		rctx, _ := context.WithTimeout(ctx, Timeout)
		if err := client.Ping(rctx, readpref.Primary()); err != nil {
			// Can't ping mongo, retry unless we exhausted our retries
			if i == limit - 1 {
				panic("exhausted our retries")
			}

			continue    // Retry
		} else {
			log.Printf("connected to %s successfully\n", uri)
			break
		}
	}
}
