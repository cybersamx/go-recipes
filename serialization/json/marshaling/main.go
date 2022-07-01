package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedUNIX int64     `json:"updated"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`              // This field will be omitted completely
	Note        string    `json:"note,omitempty"` // This field will be omitted in the output if field is empty
}

func printJSON(user User, isPrettyPrint bool) {
	var data []byte
	var err error

	if isPrettyPrint {
		prefix := ""
		indent := "  " // 2 spaces
		data, err = json.MarshalIndent(&user, prefix, indent)
	} else {
		data, err = json.Marshal(&user)
	}

	if err != nil {
		log.Fatalf("can't marshal user to data: %v", err)
	}

	fmt.Println(string(data))
}

func printUser(jsonString string) {
	data := []byte(jsonString)

	var parsedUser User
	err := json.Unmarshal(data, &parsedUser)
	if err != nil {
		log.Fatalf("can't unmarshal user to data: %v", err)
	}

	fmt.Println(parsedUser)
}

func main() {
	user := User{
		ID:          uuid.New().String(),
		CreatedAt:   time.Now(),
		UpdatedUNIX: time.Now().Unix(),
		Email:       "nobody@example.com",
		Password:    "password",
		Note:        "",
	}

	fmt.Println("--- No pretty print JSON ---")
	printJSON(user, false)
	fmt.Println("\n--- Pretty print JSON ---")
	printJSON(user, true)

	jsonString := `
		{
			"id": "5be20aa1-f6aa-4b5f-abb0-d5bfa1a76ed2",
			"created_at": "2020-01-11T23:45:37.353577-08:00",
			"updated": 1578815137,
			"email": "nobody@example.com",
			"password": "password",
			"note": "Hello"
		}
	`

	fmt.Println("\n--- Print User ---")
	printUser(jsonString)
}
