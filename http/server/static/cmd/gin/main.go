package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	port      = 8000
	staticDir = "./public"
)

func main() {
	addr := fmt.Sprintf(":%d", port)
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(staticDir, false)))
	log.Println("web server running in port", port)
	log.Fatal(router.Run(addr))
}
