package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	port    = 8000
	rootDir = "./public"
)

func parseFlag() string {
	dir := flag.String("dir", rootDir, "static html directory")
	flag.Parse()

	return *dir
}

func main() {
	dir := parseFlag()

	addr := fmt.Sprintf(":%d", port)
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(dir, false)))
	log.Println("web server running at port", port)
	log.Fatal(router.Run(addr))
}
