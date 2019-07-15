package main

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
	"time"
)

type (
	Config struct {
		GCInterval time.Duration
	}

	Entity struct {
		ID        string
		ExpiresAt time.Time
	}

	Store struct {
		Config Config
		DB     *sql.DB
		Ticker *time.Ticker
	}
)

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

func (ts *Store) Close() {
	ts.Ticker.Stop()
}

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
