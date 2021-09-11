package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	port      = 8000
	staticDir = "./public"
)

type Account struct {
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
	City    string `form:"city" binding:"required"`
	Postal  string `form:"postal" binding:"required"`
}

func main() {
	addr := fmt.Sprintf(":%d", port)
	router := gin.Default()

	router.POST("/", FormHandler())
	router.Use(static.Serve("/", static.LocalFile(staticDir, false)))

	log.Println("web server running in port", port)
	log.Fatal(router.Run(addr))
}

func FormHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var account Account
		if err := ctx.ShouldBind(&account); err != nil {
			fmt.Println(err)
			http.Redirect(ctx.Writer, ctx.Request, "/failure.html", http.StatusFound)
			return
		}

		http.Redirect(ctx.Writer, ctx.Request, "/success.html", http.StatusFound)
	}
}
