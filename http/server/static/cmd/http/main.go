package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	port    = 8000
	rootDir = "./public"
)

func parseFlag() string {
	dir := flag.String("dir", rootDir, "flush cache")
	flag.Parse()

	return *dir
}

func main() {
	dir := parseFlag()

	addr := fmt.Sprintf(":%d", port)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Println("web server running at port", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
