package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Message struct {
	ID      int
	Author  string
	Content string
}

const (
	dbFilename = "./sample.db"
)

func checkerr(err error) {
	if err != nil {
		// Note: While it's okay to call fatal in a sample code, avoid using fatal or panic in production code.
		log.Fatalln(err)
	}
}

func cleanup() {
	if _, err := os.Stat(dbFilename); !os.IsNotExist(err) {
		checkerr(os.Remove(dbFilename))
	}
}

func main() {
	cleanup()

	msgs := []Message{
		{1, "पटेल", "नमस्ते दुनिया"},
		{2, "尚氣", "你好世界"},
		{3,"Isabella", "Hola Mundo"},
		{4 ,"John", "Hello World"},
	}

	// Connect to the database.
	db, err := sql.Open("sqlite3", dbFilename)
	checkerr(err)
	defer func() {
		checkerr(db.Close())
	}()

	// Create a table.
	createStmt := `
		CREATE TABLE message (
			id INTEGER NOT NULL PRIMARY KEY,
			author TEXT,
			content TEXT
		)
	`
	_, err = db.Exec(createStmt)
	checkerr(err)

	// Insert rows.
	tx, err := db.Begin()
	checkerr(err)
	stmt, err := tx.Prepare("INSERT INTO message(id, author, content) VALUES(?, ?, ?)")
	checkerr(err)
	defer func() {
		checkerr(stmt.Close())
	}()

	for _, msg := range msgs {
		_, err = stmt.Exec(msg.ID, msg.Author, msg.Content)
		checkerr(err)
	}
	checkerr(tx.Commit())

	// Query rows.
	rows, err := db.Query("SELECT id, author, content from message")
	checkerr(err)
	defer func() {
		checkerr(db.Close())
	}()

	for rows.Next() {
		var id int
		var author, msg string
		checkerr(rows.Scan(&id, &author, &msg))
		fmt.Println(id, author, msg)
	}
	checkerr(rows.Err())

	// Update a row.
	tx, err = db.Begin()
	checkerr(err)
	stmt, err = tx.Prepare("UPDATE message SET author = ? WHERE id = ?")
	checkerr(err)
	defer func() {
		checkerr(stmt.Close())
	}()

	_, err = stmt.Exec("Javier", 3)
	checkerr(err)
	checkerr(tx.Commit())

	// Query a row.
	stmt, err = db.Prepare("SELECT id, author, content FROM message WHERE id = ?")
	checkerr(err)
	defer func() {
		checkerr(stmt.Close())
	}()

	message := Message{}
	err = stmt.QueryRow(3).Scan(&message.ID, &message.Author, &message.Content)
	checkerr(err)
	fmt.Println(message)
}
