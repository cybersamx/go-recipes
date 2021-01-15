package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Token represents token that.
type Token struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Age       int                `bson:"age,omitempty"`
	CreatedAt time.Time          `bson:"expireAt,omitempty"`
}

const (
	timeout      = 30 * time.Second
	indexTimeout = 10 * time.Second
	dbName       = "admin"
	collName     = "users"
)

func newClient(username, password, host, database string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		username,
		password,
		host,
		database)
	return mongo.NewClient(options.Client().ApplyURI(uri))
}

func connectDatabase(parentCtx context.Context, client *mongo.Client, database string) (*mongo.Database, error) {
	err := client.Connect(parentCtx)
	if err != nil {
		return nil, err
	}
	db := client.Database(database)
	return db, nil
}

func main() {
	// Using admin account to connect since some of the commands require elevated privileges.
	client, err := newClient("root", "password", "localhost", dbName)
	if err != nil {
		log.Fatalf("problem setting up a client to Mongo: %v", err)
	}

	ctx := context.Background()
	db, err := connectDatabase(ctx, client, dbName)
	if err != nil {
		log.Fatalf("problem connecting to Mongo database %s: %v", dbName, err)
	}
	defer func() {
		disCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		if err := client.Disconnect(disCtx); err != nil {
			panic(err)
		}
	}()

	// --- Operations ---

	// Create index.
	indexName, err := CreateIndex(ctx, db.Collection(collName), "email")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created index %s\n", indexName)

	// Create TTL index.
	opts := IndexOptions{isTTL: true}
	indexName, err = CreateIndex(ctx, db.Collection(collName), "createdAt", &opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created TTL index %s\n", indexName)

	// List all collections.
	collNames, err := ListCollectionNames(ctx, db)
	if err != nil {
		panic(err)
	}
	fmt.Println(collNames)

	// Does collection name exist?
	ok, err := CollectionExists(ctx, db, collName)
	if err != nil {
		panic(err)
	}
	if ok {
		fmt.Printf("Collection %s exist", collName)
	} else {
		fmt.Printf("Collection %s does not exist", collName)
	}

	// Get host info.
	hinfo, err := GetHostInfo(ctx, db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("System is running on %s OS with %d cores and %d MB of memory",
		hinfo.OS.Name,
		hinfo.System.CoresCount,
		hinfo.System.MemSize,
	)
}
