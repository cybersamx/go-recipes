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

const connectTimeout = 5 * time.Second
const pingTimeout = 2 * time.Second
const retries = 5

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Some arbitrary URI string
	uri := "mongodb://user:password@localhost:27017/database"
	client, _ := mongo.NewClient(options.Client().ApplyURI(uri))

	for i := 0; i < retries; i++ {
		// When using context functions that return CancelFunc, we must call the
		// CancelFunc the resource allocated for the context will linger and we
		// have a resource leak. But for defer, it doesn't get called until the function
		// scope ends. Since we are in a for loop block, the CancelFunc will only get
		// called only once. There are 2 solutions to this:
		// 1. Wrap the defer cancel() statement in an anonymous function.
		// 2. Don't use defer and move the cancel() all the way to the end of the loop block.
		//
		// Use option 1 to keep it timeout. Using option 2 here to illustrate the use of
		// anonymous function.
		ok := func() bool {
			log.Printf("attempt %d: connecting to mongo %s\n", i + 1, uri)
			cctx, ccancel := context.WithTimeout(ctx, connectTimeout)
			defer ccancel()
			err := client.Connect(cctx)
			// client.Disconnect() returns an error, which should be checked. To handle
			//  an error with defer, you need to wrap the call in an anonymous function
			defer func() {
				err := client.Disconnect(cctx)
				if err != nil {
					log.Panicf("can't close connection to mongo: %v\n", err)
				}
			}()

			if err != nil && err != topology.ErrTopologyConnected {
				log.Printf("can't connect to mongo: %v\n", err)
				return false
			}

			pctx, pcancel := context.WithTimeout(ctx, pingTimeout)
			defer pcancel()
			if err := client.Ping(pctx, readpref.Primary()); err != nil {
				// Ping mongo failed, retry unless we exhausted our retries
				if i == retries - 1 {
					log.Printf("exhausted our retries\n")
					return true
				}

				return false
			}

			log.Print("connected to mongo successfully\n")
			return true
		}()

		if ok {
			break
		}
	}
}
