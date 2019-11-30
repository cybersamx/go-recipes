package main

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

type config struct {
    host string
    database string
    username string
    password string
}

func newDefaultConfig() config {
    // NOTE: Do not do the following in your code. This is done for just for the sample code.
    // 1. Do not connect to Mongo as the global root user.
    // 2. Do not hard code passwords or other sensitive data.
    return config{
        host:     "localhost",
        database: "test",
        username: "root",
        password: "password",
    }
}

func newMongoDatabase(cfg config, ctx context.Context) (*mongo.Database, error) {
    uri := fmt.Sprintf("mongodb://%s:%s@%s/%s",
        cfg.username,
        cfg.password,
        cfg.host,
        cfg.database)
    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        return nil, fmt.Errorf("can't connect to mongo. %v", err)
    }
    err = client.Connect(ctx)
    if err != nil {
        return nil, fmt.Errorf("can't connect to mongo. %v", err)
    }

    db := client.Database(cfg.database)
    return db, nil
}

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    db, err := newMongoDatabase(newDefaultConfig(), ctx)
    if err != nil {
        log.Fatalf("Can't open database")
    }
    fmt.Printf("Connected: %v", db)
}
