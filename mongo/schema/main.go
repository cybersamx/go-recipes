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

// User represents a user account of an application.
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username,omitempty"`
	Email      string             `bson:"email,omitempty"`
	City       string             `bson:"city,omitempty"`
	Age        int                `bson:"age,omitempty"`
	CreatedAt  time.Time          `bson:"createdAt,omitempty"`
	ModifiedAt time.Time          `bson:"modifiedAt,omitempty"`
}

const (
	timeout  = 30 * time.Second
	dbName   = "go-recipes"
	collName = "users"
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

func createUser(parentCxt context.Context, db *mongo.Database, user User) error {
	ctx, cancel := context.WithTimeout(parentCxt, timeout)
	defer cancel()

	_, err := db.Collection(collName).InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("problem inserting a user: %v", err)
	}

	return nil
}

func updateUser(parentCtx context.Context, db *mongo.Database, user User) (int64, error) {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	result, err := db.Collection(collName).UpdateOne(
		ctx,
		bson.D{{"_id", user.ID}},
		bson.D{{"$set", user}},
	)
	if err != nil {
		return 0, err
	}

	return result.UpsertedCount, err
}

func deleteUser(parentCtx context.Context, db *mongo.Database, userID primitive.ObjectID) (int64, error) {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	result, err := db.Collection(collName).DeleteOne(ctx, bson.D{
		{"_id", userID},
	})
	if err != nil {
		return 0, fmt.Errorf("problem deleting id %s from database", userID)
	}

	return result.DeletedCount, nil
}

func getUser(parentCtx context.Context, db *mongo.Database, userID primitive.ObjectID) (*User, error) {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	var result User
	err := db.Collection(collName).FindOne(ctx, bson.D{
		{"_id", userID},
	}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func printUser(user *User) {
	fmt.Println(user)
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

	// Create user
	user := User{
		ID:         primitive.NewObjectID(),
		Username:   "tommy",
		Email:      "tommy@example.com",
		City:       "Singapore",
		Age:        30,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	if err := createUser(ctx, db, user); err != nil {
		log.Fatalf("problem creating a user: %v", err)
	}

	fmt.Printf("successfully created user %s, userID %s\n", user.Username, user.ID)

	// Get user
	foundUser, err := getUser(ctx, db, user.ID)
	if err != nil {
		log.Fatalf("problem finding user %s", user.ID)
	}

	fmt.Printf("successfully found user %s, userID %s\n", foundUser.Username, user.ID)
	printUser(foundUser)

	// Update user
	user.City = "London"
	upsertCount, err := updateUser(ctx, db, user)
	if err != nil {
		log.Fatalf("problem updating user %s", user.ID)
	}

	fmt.Printf("successfully updated %d user %s, userID %s\n", upsertCount, user.Username, user.ID)

	// Get user
	foundUser, err = getUser(ctx, db, user.ID)
	if err != nil {
		log.Fatalf("problem finding user %s", user.ID)
	}

	fmt.Printf("successfully found user %s, userID %s\n", foundUser.Username, user.ID)
	printUser(foundUser)

	// Delete user
	deleteCount, err := deleteUser(ctx, db, user.ID)
	if err != nil {
		log.Fatalf("problem deleting a user: %v", err)
	}
	fmt.Printf("successfully deleted %d, userID %s\n", deleteCount, user.ID)
}
