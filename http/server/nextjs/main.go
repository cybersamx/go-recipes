package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"syreclabs.com/go/faker"
)

//go:embed web/dist
//go:embed web/dist/_next
//go:embed web/dist/_next/static/*/*.js
//go:embed web/dist/_next/static/chunks/pages/*.js
var webFS embed.FS

const addr = ":3000"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(faker.Name().Name()))
}

func main() {
	rootFS, err := fs.Sub(webFS, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	// Content on the web/dist folder will be served as root /.
	http.Handle("/", http.FileServer(http.FS(rootFS)))
	//// The API will be served under `/api`.
	http.HandleFunc("/api", apiHandler)

	log.Printf("starting http server at %v", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
