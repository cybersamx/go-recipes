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
	ID         string
	Username   string
	Email      string
	City       string
	Age        int
	CreatedAt  time.Time
	ModifiedAt time.Time
}

const (
	timeout = 30 * time.Second
	dbName  = "go-recipes"
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

func createUser(db *mongo.Database, user User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result, err := db.Collection("users").InsertOne(ctx, bson.D{
		{"username", user.Username},
		{"email", user.Email},
		{"city", user.City},
		{"age", user.Age},
		{"createdAt", user.CreatedAt},
		{"modifiedAt", user.ModifiedAt},
	})
	if err != nil {
		return "", fmt.Errorf("problem inserting a user: %v", err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func updateUser(db *mongo.Database, user User) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return 0, fmt.Errorf("problem converting a string ID to an ObjectID: %v", err)
	}
	result, err := db.Collection("users").UpdateOne(
		ctx,
		bson.D{
			{"_id", objectID},
		},
		bson.D{
			{"$set", bson.D{
				{"username", user.Username},
				{"email", user.Email},
				{"city", user.City},
				{"age", user.Age},
				{"modifiedAt", user.ModifiedAt},
			}}})
	if err != nil {
		return 0, err
	}

	return result.UpsertedCount, err
}

func deleteUser(db *mongo.Database, userID string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return 0, fmt.Errorf("problem converting a string ID to an ObjectID: %v", err)
	}
	result, err := db.Collection("users").DeleteOne(ctx, bson.D{
		{"_id", objectID},
	})
	if err != nil {
		return 0, fmt.Errorf("problem deleting id %s from database", userID)
	}

	return result.DeletedCount, nil
}

func getUser(db *mongo.Database, userID string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("problem converting a string ID to an ObjectID: %v", err)
	}

	var result User
	err = db.Collection("users").FindOne(ctx, bson.D{
		{"_id", objectID},
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

	db, err := connectDatabase(context.Background(), client, dbName)
	if err != nil {
		log.Fatalf("problem connecting to Mongo database %s: %v", dbName, err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Create user
	user := User{
		Username:   "tommy",
		Email:      "tommy@example.com",
		City:       "Singapore",
		Age:        30,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	id, err := createUser(db, user)
	if err != nil {
		log.Fatalf("problem creating a user: %v", err)
	}

	user.ID = id
	fmt.Printf("successfully created user %s, userID %s\n", user.Username, user.ID)

	// Get user
	foundUser, err := getUser(db, user.ID)
	if err != nil {
		log.Fatalf("problem finding user %s", user.ID)
	}

	fmt.Printf("successfully found user %s, userID %s\n", foundUser.Username, user.ID)
	printUser(foundUser)

	// Update user
	user.City = "London"
	upsertCount, err := updateUser(db, user)
	if err != nil {
		log.Fatalf("problem updating user %s", user.ID)
	}

	fmt.Printf("successfully updated %d user %s, userID %s\n", upsertCount, user.Username, user.ID)

	// Get user
	foundUser, err = getUser(db, user.ID)
	if err != nil {
		log.Fatalf("problem finding user %s", user.ID)
	}

	fmt.Printf("successfully found user %s, userID %s\n", foundUser.Username, user.ID)
	printUser(foundUser)

	// Delete user
	deleteCount, err := deleteUser(db, user.ID)
	if err != nil {
		log.Fatalf("problem deleting a user: %v", err)
	}
	fmt.Printf("successfully deleted %d, userID %s\n", deleteCount, user.ID)
}
