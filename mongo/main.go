package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type config struct {
	host     string
	database string
	username string
	password string
}

func newDefaultConfig() config {
	// NOTE: Do not do the following in your code. This is done for just for the sample code.
	// Do not hard code passwords or other sensitive data.
	return config{
		host:     "localhost",
		database: "go-recipes",
		username: "nobody",
		password: "secrets",
	}
}

func newClient(cfg config) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		cfg.username,
		cfg.password,
		cfg.host,
		cfg.database)
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

func createUser(ctx context.Context, db *mongo.Database, user User) (string, error) {
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

func updateUser(ctx context.Context, db *mongo.Database, user User) (int64, error) {
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

func deleteUser(ctx context.Context, db *mongo.Database, userID string) (int64, error) {
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

func getUser(ctx context.Context, db *mongo.Database, userID string) (*User, error) {
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
	cfg := newDefaultConfig()
	client, err := newClient(cfg)
	if err != nil {
		log.Fatalf("problem setting up a client to Mongo: %v", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	db, err := connectDatabase(ctx, client, cfg.database)
	if err != nil {
		log.Fatalf("problem connecting to Mongo database %s: %v", cfg.database, err)
		os.Exit(1)
	}

	// Create user
	user := User{
		Username:   "tommy",
		Email:      "tommy@example.com",
		City:       "Singapore",
		Age:        30,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	id, err := createUser(ctx, db, user)
	if err != nil {
		log.Fatalf("problem creating a user: %v", err)
		os.Exit(1)
	}

	user.ID = id
	fmt.Printf("successfully created user %s, userID %s\n", user.Username, user.ID)

	// Get user
	foundUser, err := getUser(ctx, db, user.ID)
	if err != nil {
		log.Fatalf("problem finding user %s", user.ID)
		os.Exit(1)
	}

	fmt.Printf("successfully found user %s, userID %s\n", foundUser.Username, user.ID)
	printUser(foundUser)

	// Update user
	user.City = "London"
	upsertCount, err := updateUser(ctx, db, user)
	if err != nil {
		log.Fatalf("problem updating user %s", user.ID)
		os.Exit(1)
	}

	fmt.Printf("successfully updated %d user %s, userID %s\n", upsertCount, user.Username, user.ID)

	// Get user
	foundUser, err = getUser(ctx, db, user.ID)
	if err != nil {
		log.Fatalf("problem finding user %s", user.ID)
		os.Exit(1)
	}

	fmt.Printf("successfully found user %s, userID %s\n", foundUser.Username, user.ID)
	printUser(foundUser)

	// Delete user
	deleteCount, err := deleteUser(ctx, db, user.ID)
	if err != nil {
		log.Fatalf("problem deleting a user: %v", err)
		os.Exit(1)
	}
	fmt.Printf("successfully deleted %d, userID %s\n", deleteCount, user.ID)
}
