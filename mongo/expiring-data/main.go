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
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Value    string             `bson:"value,omitempty"`
	ExpireAt time.Time          `bson:"expireAt,omitempty"`
}

const (
	timeout      = 30 * time.Second
	ttl          = 60 * time.Second
	indexTimeout = 10 * time.Second
	collName  = "tokens"
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

func createToken(parentCtx context.Context, db *mongo.Database, token Token) error {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	_, err := db.Collection("tokens").InsertOne(ctx, token)
	if err != nil {
		return fmt.Errorf("problem inserting a user: %v", err)
	}

	return nil
}

func containsKey(m primitive.M, key string) bool {
	_, ok := m[key]
	return ok
}

// containsChildKey returns true if param key matches the child key in { "key": { key: "someValue" }}
func containsChildKey(m primitive.M, key string) bool {
	val := m["key"]
	if val != nil {
		childMap := val.(primitive.M)
		_, ok := childMap[key]
		return ok
	}

	return false
}

// ttlIndexExist checks to see if a TTL index for `field` in `collection` already exist.
func ttlIndexExist(parentCtx context.Context, collection *mongo.Collection, field string) (bool, error) {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	cursor, err := collection.Indexes().List(ctx)
	if err != nil {
		return false, err
	}

	var indexes []bson.M
	if err := cursor.All(ctx, &indexes); err != nil {
		return false, err
	}

	for _, m := range indexes {
		if containsKey(m, "expireAfterSeconds") && containsChildKey(m, field) {
			return true, nil
		}
	}

	return false, nil
}

func createTTLIndex(parentCtx context.Context, collection *mongo.Collection, field string) error {
	ok, err := ttlIndexExist(parentCtx, collection, field)
	if err != nil {
		return err
	}
	if ok {
		// Skip index creation given the index already exist.
		return nil
	}

	// Create TTL index
	models := []mongo.IndexModel{
		{
			Keys: bson.D{
				{"expireAt", 1},
			},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
	}

	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	opts := options.CreateIndexes().SetMaxTime(indexTimeout)
	names, err := collection.Indexes().CreateMany(ctx, models, opts)

	fmt.Println("created indexes", names)

	return err
}

func main() {
	client, err := newClient("nobody", "secrets", "localhost", dbName)
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

	if err := createTTLIndex(ctx, db.Collection(collName), "expireAt"); err != nil {
		panic(err)
	}

	token := Token{
		ID:       primitive.NewObjectID(),
		Value:    "SomeSecretValue",
		ExpireAt: time.Now().Add(ttl),
	}

	if err := createToken(ctx, db, token); err != nil {
		log.Fatalf("problem creating a user: %v", err)
	}
}
