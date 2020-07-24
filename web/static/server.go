package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port    = 8000
	rootDir = "./root"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./root")))
	log.Println("web server running at port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
