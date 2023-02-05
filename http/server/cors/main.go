package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//go:embed public
var pub embed.FS

const (
	webSrvAddr   = ":8000"
	imageSrvAddr = ":9000"
)

func startGinImageServer(ctx context.Context, enableCORS bool) {
	router := gin.Default()

	// Middleware for handling cors requests.
	if enableCORS {
		cfg := cors.Config{
			AllowAllOrigins: false,
			AllowOrigins:    []string{"http://localhost:8000"},
		}

		router.Use(cors.New(cfg))
	}

	// Handler to handle incoming requests to retrieve the image.
	router.GET("/image", func(ctx *gin.Context) {
		// Log the header. If the caller is from a web browser, we should see the Origin header.
		log.Printf("Received Origin=%s", ctx.Request.Header.Get("Origin"))

		// Response back to the caller.
		buf, err := pub.ReadFile("public/sample.svg")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.Writer.Header().Set("Content-Type", "image/svg+xml")
		ctx.Writer.Header().Set("Content-Length", strconv.Itoa(len(buf)))

		ctx.Writer.Write(buf)
	})

	log.Println("image server running at addr", imageSrvAddr)
	err := router.Run(imageSrvAddr)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("image server shutting down")
	} else if err != nil {
		log.Fatalln(err)
	}
}

func abortError(w http.ResponseWriter, code int) {
	var message string

	switch code {
	case http.StatusMethodNotAllowed:
		message = "method not allowed"
	case http.StatusInternalServerError:
		message = "internal server error"
	case http.StatusForbidden:
		message = "forbidden"
	}

	w.WriteHeader(code)
	n, _ := fmt.Fprintf(w, "%d %s", code, message)
	w.Header().Set("Content-Length", strconv.Itoa(n))
}

func startGorillaImageServer(ctx context.Context, enableCORS bool) {
	router := mux.NewRouter()

	// Handler to handle incoming requests to retrieve the image.
	imageHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		if r.Method != "GET" {
			abortError(w, http.StatusMethodNotAllowed)
			return
		}

		// Log the header. If the caller is from a web browser, we should see the Origin header.
		log.Printf("Received Origin=%s", r.Header.Get("Origin"))

		// Response back to the caller.
		buf, err := pub.ReadFile("public/sample.svg")
		if err != nil {
			abortError(w, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
		w.Write(buf)
	}

	var handler http.Handler

	// Middleware for handling cors requests.
	if enableCORS {
		gcors := handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:8000"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Content-Length", "Origin"}),
		)

		router.HandleFunc("/image", imageHandler)
		handler = gcors(router)
	} else {
		router.HandleFunc("/image", imageHandler)
		handler = router
	}

	log.Println("image server running at addr", imageSrvAddr)
	err := http.ListenAndServe(imageSrvAddr, handler)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("image server shutting down")
	} else if err != nil {
		log.Fatalln(err)
	}
}

func startGinWebServer(ctx context.Context) {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		buf, err := pub.ReadFile("public/index.html")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.Status(http.StatusOK)
		ctx.Writer.Write(buf)

		ctx.Writer.Header().Set("Content-Type", "text/html")
		ctx.Writer.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	})

	log.Println("web server running at addr", webSrvAddr)
	err := router.Run(webSrvAddr)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("web server shutting down")
	} else if err != nil {
		log.Fatalln(err)
	}
}

func startGorillaWebServer(ctx context.Context) {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			abortError(w, http.StatusMethodNotAllowed)
			return
		}

		buf, err := pub.ReadFile("public/index.html")
		if err != nil {
			abortError(w, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(buf)

		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	})

	log.Println("web server running at addr", webSrvAddr)
	server := http.Server{
		Addr:    webSrvAddr,
		Handler: router,
	}
	defer server.Close()

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("web server shutting down")
	} else if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	enableCORS := flag.Bool("cors", false, "enable cors support")
	framework := flag.String("framework", "gin", "which framework to use (gin|gorilla)")

	flag.Parse()

	if *framework != "gin" && *framework != "gorilla" {
		fmt.Println("flag framework must be either gin or gorilla")
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if *framework == "gorilla" {
			startGorillaImageServer(ctx, *enableCORS)
			return
		}

		startGinImageServer(ctx, *enableCORS)
	}()

	go func() {
		if *framework == "gorilla" {
			startGorillaWebServer(ctx)
			return
		}

		startGinWebServer(ctx)
	}()

	<-ctx.Done()
}
