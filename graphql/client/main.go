package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/machinebox/graphql"
)

type Language struct {
	Name string
}

type Country struct {
	Code string
	Name string
	Languages []Language
}

type ResponseData struct {
	Countries []Country
}

func main() {
	// --- Simple query without filters ---

	client := graphql.NewClient("https://countries.trevorblades.com/")

	req := graphql.NewRequest(`
		query {
            countries {
                code,
                name,
                languages {
                    name
                }
			}
		}
	`)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var res ResponseData
	if err := client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}

	fmt.Println("GraphQL query without filters")
	fmt.Println(res)

	// --- Simple query without filters ---

	req = graphql.NewRequest(`
		query ($key: String!) {
			countries(filter: {
                code: { eq: $key }
            }) {
                code,
                name,
                languages {
                    name
                }
			}
		}
	`)

	country := "US"
	req.Var("key", country)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if err := client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("GraphQL query matching country =", country)
	fmt.Println(res)
}
