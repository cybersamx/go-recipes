package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port      = 8000
	staticDir = "./public"
)

func main() {
	addr := fmt.Sprintf(":%d", port)
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	log.Println("web server running in port", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
