package main

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"time"
)

type (
	// Config encapsulates the configurations for this application.
	Config struct {
		GCInterval time.Duration
	}

	// Entity is a timeout entity for the SQL database.
	Entity struct {
		ID        string
		ExpiresAt time.Time
	}

	// Store encapsulates the connection and configurations to the SQL database.
	Store struct {
		Config Config
		DB     *sql.DB
		Ticker *time.Ticker
	}
)

// NewStore creates a new instance of Store.
func NewStore(db *sql.DB, config Config) (*Store, error) {
	ts := Store{
		Config: config,
		DB:     db,
		Ticker: time.NewTicker(config.GCInterval),
	}

	_, err := ts.DB.Exec(`
CREATE TABLE IF NOT EXISTS Entities (
		    id TEXT PRIMARY KEY,
		    expire_at DATETIME
		)`)
	if err != nil {
		return nil, err
	}

	ts.Ticker = time.NewTicker(ts.Config.GCInterval)
	go ts.garbageCollect()

	return &ts, err
}

func (ts *Store) garbageCollect() {
	for range ts.Ticker.C {
		now := time.Now()
		result, err := ts.DB.Exec(`
			DELETE FROM Entities WHERE
			expire_at <= ?`, now)
		if err != nil {
			log.Fatal(err)
		}
		count, _ := result.RowsAffected()
		log.Printf("garbage collection called, removed %d entities that are older than %v\n", count, now)
	}
}

// Close close the connection to the SQL database and clean up all allocated resources
// associated with the SQL database connection.
func (ts *Store) Close() {
	ts.Ticker.Stop()
}

// Create creates a new entity in the database.
func (ts *Store) Create(expiresAt time.Time) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	idString := id.String()
	_, err = ts.DB.Exec(`
		INSERT INTO Entities (id, expire_at)
    	VALUES (?, ?)`, idString, expiresAt)
	if err != nil {
		return "", err
	}

	return idString, nil
}

// Get retrieves the entity from the SQL database.
func (ts *Store) Get(id string) (*Entity, error) {
	rows, err := ts.DB.Query(`
		SELECT expire_at FROM Entities
		WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	var expireAt time.Time
	if rows.Next() {
		rows.Scan(&expireAt)
		return &Entity{
			ID:        id,
			ExpiresAt: expireAt,
		}, nil
	}

	return nil, nil
}
