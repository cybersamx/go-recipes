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

type config struct {
	host     string
	database string
	username string
	password string
}

// User represents a user account of an application.
type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Username   string             `bson:"username"`
	Email      string             `bson:"email"`
	City       string             `bson:"city"`
	Age        int                `bson:"age"`
	CreatedAt  time.Time          `bson:"createdAt"`
	ModifiedAt time.Time          `bson:"modifiedAt"`
}


func newDefaultConfig() config {
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

func createUser(ctx context.Context, db *mongo.Database, user User) error {
	_, err := db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("problem inserting a user: %v", err)
	}

	return nil
}

func updateUser(ctx context.Context, db *mongo.Database, user User) (int64, error) {
	result, err := db.Collection("users").UpdateOne(
		ctx,
		bson.D{{"_id", user.ID}},
		bson.D{{"$set", user}},
	)
	if err != nil {
		return 0, err
	}

	return result.UpsertedCount, err
}

func deleteUser(ctx context.Context, db *mongo.Database, userID primitive.ObjectID) (int64, error) {
	result, err := db.Collection("users").DeleteOne(ctx, bson.D{
		{"_id", userID},
	})
	if err != nil {
		return 0, fmt.Errorf("problem deleting id %s from database", userID)
	}

	return result.DeletedCount, nil
}

func getUser(ctx context.Context, db *mongo.Database, userID primitive.ObjectID) (*User, error) {
	var result User
	err := db.Collection("users").FindOne(ctx, bson.D{
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
	cfg := newDefaultConfig()
	client, err := newClient(cfg)
	if err != nil {
		log.Fatalf("problem setting up a client to Mongo: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	db, err := connectDatabase(ctx, client, cfg.database)
	if err != nil {
		log.Fatalf("problem connecting to Mongo database %s: %v", cfg.database, err)
	}
	defer client.Disconnect(ctx)

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
