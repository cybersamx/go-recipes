package main

import (
	"context"
	"github.com/flowchartsman/retry"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
	"log"
	"time"
)

const retries = 5
const connectTimeout = 5 * time.Second
const pingTimeout = 2 * time.Second
const initialDelay = 1 * time.Second
const maxDelay = 6 * time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Some arbitrary URI string
	uri := "mongodb://user:password@localhost:27017/database"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("can't create an instance of mongo client")
	}

	retrier := retry.NewRetrier(retries, initialDelay, maxDelay)
	err = retrier.RunContext(ctx, func(ctx context.Context) (retErr error) {
		log.Printf("trying to connect to mongo %s\n", uri)
		cctx, ccancel := context.WithTimeout(ctx, connectTimeout)
		defer ccancel()
		err := client.Connect(cctx)
		// client.Disconnect() returns an error, which should be checked. To handle
		// an error with defer, you need to wrap the call in an anonymous function
		// and we use a named return param retErr to return the error.
		defer func() {
			err := client.Disconnect(cctx)
			if err != nil {
				log.Printf("can't close connection to mongo: %v\n", err)
				retErr = err
			}
		}()

		if err != nil && err != topology.ErrTopologyConnected {
			log.Printf("can't connect to mongo: %v\n", err)
			retErr = err
			return // Return the err to retrier telling it to retry.
		}

		pctx, pcancel := context.WithTimeout(ctx, pingTimeout)
		defer pcancel()
		if err := client.Ping(pctx, readpref.Primary()); err != nil {
			retErr = err
			return // Return the err to retrier telling it to retry.
		}

		return nil // Success
	})

	if err != nil {
		log.Printf("exhausted our retries\n")
		return
	}

	log.Print("connected to mongo successfully\n")
}
