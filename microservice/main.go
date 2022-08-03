package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cybersamx/go-recipes/microservice/simple/endpoints"
	"github.com/cybersamx/go-recipes/microservice/simple/services"
	"github.com/cybersamx/go-recipes/microservice/simple/transports"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

const (
	addr = "localhost:8000"
)

func main() {
	// Build http handlers.
	svc := services.StateServiceImpl{}

	getStatesHandler := httptransport.NewServer(
		endpoints.GetStatesEndpoint(svc),
		transports.DecodeGetStatesRequest(),
		httptransport.EncodeJSONResponse,
	)

	getStateHandler := httptransport.NewServer(
		endpoints.GetStateEndpoint(svc),
		transports.DecodeGetStateRequest(),
		httptransport.EncodeJSONResponse,
	)

	// Set up HTTP server.
	router := gin.Default()
	router.GET("/states", func(gtx *gin.Context) {
		getStatesHandler.ServeHTTP(gtx.Writer, gtx.Request)
	})
	router.GET("/states/:abbreviation", func(gtx *gin.Context) {
		getStateHandler.ServeHTTP(gtx.Writer, gtx.Request)
	})

	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Run the server and set up for a graceful shutdown.
	errs := make(chan error)

	go func() {
		log.Printf("starting http server %s", addr)
		errs <- srv.ListenAndServe()
	}()

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("received signal: %s", <-quit)
	}()

	log.Printf("server exiting due to %s", <-errs)
}
