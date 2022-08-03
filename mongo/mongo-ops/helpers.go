package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// HostInfo maps to the schema of the result that return from running mongo command
// runCommand({ hostInfo: 1 }).
// See https://docs.mongodb.com/manual/reference/command/hostInfo/
type HostInfo struct {
	System struct {
		Hostname   string `bson:"hostname,omitempty"`
		CoresCount int    `bson:"numCores,omitempty"`
		MemSize    int    `bson:"memSizeMB,omitempty"`
	}
	OS struct {
		Type    string `bson:"type,omitempty"`
		Name    string `bson:"name,omitempty"`
		Version string `bson:"version,omitempty"`
	}
}

type IndexOptions struct {
	isTTL bool
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

func containsString(texts []string, val string) bool {
	for _, text := range texts {
		if text == val {
			return true
		}
	}

	return false
}

// IndexExists checks to see if index for `field` in `collection` already exist.
func IndexExists(parentCtx context.Context, collection *mongo.Collection, field string, opts ...*IndexOptions) (bool, error) {
	var isTTL bool
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		isTTL = opt.isTTL
	}

	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	cursor, err := collection.Indexes().List(ctx)
	if err != nil {
		return false, err
	}

	var indexes []bson.M
	if err := cursor.All(context.Background(), &indexes); err != nil {
		return false, err
	}

	for _, m := range indexes {
		if isTTL {
			if containsKey(m, "expireAfterSeconds") && containsChildKey(m, field) {
				return true, nil
			}
		} else {
			if containsChildKey(m, field) {
				return true, nil
			}
		}
	}

	return false, nil
}

// CreateIndex creates an index associated with `field` in `collection`.
func CreateIndex(parentCtx context.Context, collection *mongo.Collection, field string, opts ...*IndexOptions) (string, error) {
	var isTTL bool
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		isTTL = opt.isTTL
	}

	ok, err := IndexExists(parentCtx, collection, field, opts...)
	if err != nil {
		return "", err
	}
	if ok {
		// Skip index creation given the index already exist.
		return "", nil
	}

	// Create TTL index
	model := mongo.IndexModel{
		Keys: bson.D{
			{field, 1},
		},
	}

	if isTTL {
		// TODO: We may want to check that the field is of type Date as it is needed for TTL to work properly.
		model.Options = options.Index().SetExpireAfterSeconds(0)
	}

	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	ciOpts := options.CreateIndexes().SetMaxTime(indexTimeout)
	indexName, err := collection.Indexes().CreateOne(ctx, model, ciOpts)

	return indexName, err
}

// ListCollections returns all collections in dataabase `db`.
func ListCollectionNames(parentCtx context.Context, db *mongo.Database) ([]string, error) {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	return db.ListCollectionNames(ctx, bson.D{})
}

// CollectionExists return true if collection named `collName` exist in database `db`.
func CollectionExists(parentCtx context.Context, db *mongo.Database, collName string) (bool, error) {
	collNames, err := ListCollectionNames(parentCtx, db)
	if err != nil {
		return false, err
	}

	return containsString(collNames, collName), nil
}

// GetHostInfo returns info about the hosting system. It's similar to running the following in the mongo shell:
// `db.runCommand({ hostInfo: 1 })
// For more info on the underlying command, see https://docs.mongodb.com/manual/reference/command/hostInfo/
func GetHostInfo(parentCtx context.Context, db *mongo.Database) (*HostInfo, error) {
	cmd := bson.D{
		{"hostInfo", 1},
	}

	var hostInfo HostInfo
	if err := db.RunCommand(parentCtx, cmd).Decode(&hostInfo); err != nil {
		return nil, err
	}

	return &hostInfo, nil
}
