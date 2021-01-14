package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Token represents token that.
type Token struct {
	ID       primitive.ObjectID `bson:"_id"`
	Value    string             `bson:"value"`
	ExpireAt time.Time          `bson:"expireAt"`
}

const (
	timeout      = 30 * time.Second
	ttl          = 60 * time.Second
	indexTimeout = 10 * time.Second
	collection   = "tokens"
	dbName       = "go-recipes"
)

func newClient(username, password, host, database string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		username,
		password,
		host,
		database)
	return mongo.NewClient(options.Client().ApplyURI(uri))
}

func connectDatabase(ctx context.Context, client *mongo.Client, database string) (*mongo.Database, error) {
	err := client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	db := client.Database(database)
	return db, nil
}

func dropCollection(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return db.Drop(ctx)
}

func createToken(db *mongo.Database, token Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := db.Collection("tokens").InsertOne(ctx, token)
	if err != nil {
		return fmt.Errorf("problem inserting a user: %v", err)
	}

	return nil
}

func createTTLIndex(db *mongo.Database) error {
	// Create TTL index
	models := []mongo.IndexModel{
		{
			Keys: bson.D{
				{"expireAt", 1},
			},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
	}

	opts := options.CreateIndexes().SetMaxTime(indexTimeout)
	names, err := db.Collection(collection).Indexes().CreateMany(context.Background(), models, opts)

	fmt.Println("created indexes", names)

	return err
}

func main() {
	client, err := newClient("nobody", "secrets", "localhost", dbName)
	if err != nil {
		log.Fatalf("problem setting up a client to Mongo: %v", err)
	}

	db, err := connectDatabase(context.Background(), client, dbName)
	if err != nil {
		log.Fatalf("problem connecting to Mongo database %s: %v", dbName, err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	if err := dropCollection(db); err != nil {
		panic(err)
	}

	if err := createTTLIndex(db); err != nil {
		panic(err)
	}

	token := Token{
		ID:       primitive.NewObjectID(),
		Value:    "SomeSecretValue",
		ExpireAt: time.Now().Add(ttl),
	}

	if err := createToken(db, token); err != nil {
		log.Fatalf("problem creating a user: %v", err)
	}
}
