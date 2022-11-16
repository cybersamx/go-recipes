package main

import (
	"fmt"
	"net/url"
)

func main() {
	apiURL, err := url.Parse("https://localhost:8000/api/v3/location/3a/shelf/5ff")
	if err != nil {
		panic(err)
	}

	// This won't work.
	apiURL.Query().Set("sort", "id")
	apiURL.Query().Set("index", "3")
	apiURL.Query().Set("page_size", "10")

	// Output: https://localhost:8000/api/v3/location/3a/shelf/5ff
	fmt.Println(apiURL)

	query := apiURL.Query()
	query.Set("sort", "id")
	query.Set("index", "3")
	query.Set("page_size", "10")
	query.Add("sort", "mod_time")

	// Set the RawQuery directly with url-encoded query parameters.
	apiURL.RawQuery = query.Encode()

	// Output: `https://localhost:8000/api/v3/location/3a/shelf/5ff?index=3&page_size=10&sort=id&sort=mod_time
	fmt.Println(apiURL)
}
