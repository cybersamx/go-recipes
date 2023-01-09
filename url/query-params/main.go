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

	// This won't work. Need to run Encode to get the encoded query params string.
	apiURL.Query().Set("sort", "id")
	apiURL.Query().Set("index", "3")
	apiURL.Query().Set("page_size", "10")

	// Output: https://localhost:8000/api/v3/location/3a/shelf/5ff
	fmt.Println(apiURL)

	// A key in query params can have multiple values. In the following, the sort key has
	// values `id` and `mod_time`.
	query := apiURL.Query()
	query.Set("sort", "id")
	query.Set("index", "3")
	query.Set("page_size", "10")
	query.Add("sort", "mod_time")

	// Set the RawQuery directly with url-encoded query parameters.
	apiURL.RawQuery = query.Encode()

	// Output: `https://localhost:8000/api/v3/location/3a/shelf/5ff?index=3&page_size=10&sort=id&sort=mod_time
	fmt.Println(apiURL)

	// Decoding just the query string, ie. index=3&page_size=10&sort=id&sort=mod_time
	fmt.Println(apiURL.RawQuery)
	params, err := url.ParseQuery(apiURL.RawQuery)
	if err != nil {
		panic(err)
	}

	for k, v := range params {
		fmt.Println(k, v)
	}
}
