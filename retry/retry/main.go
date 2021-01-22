package main

import (
	"context"
	"log"
	"time"

	"github.com/flowchartsman/retry"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

const (
	retries      = 5
	initialDelay = 1 * time.Second
	maxDelay     = 6 * time.Second
)

func newClient(parent context.Context, uri string, retries int, initialDelay, maxDelay time.Duration) (*mongo.Client, error) {
	var client *mongo.Client
	timeout := maxDelay

	retrier := retry.NewRetrier(retries, initialDelay, maxDelay)
	err := retrier.RunContext(parent, func(ctx context.Context) (retErr error) {
		var cterr error
		client, cterr = mongo.NewClient(options.Client().ApplyURI(uri))
		if cterr != nil {
			log.Fatal("can't create an instance of mongo client")
		}

		// Disconnect only if we can't connect or ping the datastore.
		//
		// client.Disconnect() returns an error, which should be checked. We wrap the function call
		// in an anonymous function so that we can capture the error.
		closeFn := func() {
			log.Printf("disconnect from mongo")
			dctx, dcancel := context.WithTimeout(ctx, timeout)
			defer dcancel()
			if derr := client.Disconnect(dctx); derr != nil && derr != mongo.ErrClientDisconnected {
				log.Printf("can't close connection to mongo: %v\n", derr)
				retErr = derr
			}
		}

		// Connect the datastore.
		log.Printf("attempting to connect mongo %s\n", uri)
		cctx, ccancel := context.WithTimeout(ctx, timeout)
		defer ccancel()
		if cerr := client.Connect(cctx); cerr != nil {
			defer closeFn()

			if cerr == topology.ErrTopologyConnected {
				// Already connected, so continue to ping the datastore.
			} else {
				log.Printf("can't connect mongo: %v\n", cerr)
				retErr = cerr
				return // Return the err to retrier telling it to retry.
			}
		}

		// Ping the datastore.
		log.Printf("attempting to ping mongo %s\n", uri)
		pctx, pcancel := context.WithTimeout(ctx, timeout)
		defer pcancel()
		if perr := client.Ping(pctx, readpref.Primary()); perr != nil {
			defer closeFn()
			log.Printf("can't ping mongo: %v\n", perr)
			retErr = perr
			return // Return the err to retrier telling it to retry.
		}

		retErr = nil
		return
	})

	return client, err
}

func main() {
	uri := "mongodb://nobody:secrets@localhost:27017/go-recipes"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := newClient(ctx, uri, retries, initialDelay, maxDelay)

	defer func() {
		log.Printf("final disconnect from mongo")
		dctx, dcancel := context.WithTimeout(ctx, maxDelay)
		defer dcancel()
		if derr := client.Disconnect(dctx); derr != nil && derr != mongo.ErrClientDisconnected {
			panic(derr)
		}
	}()

	if err != nil {
		log.Printf("exhausted our retries\n")
		return
	}

	log.Print("connected to mongo successfully\n")
}
