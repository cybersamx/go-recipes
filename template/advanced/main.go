package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	port         = 8000
	templateFile = "public/index.html"
	successURI   = "/success.html"
)

func renderTemplate(w http.ResponseWriter, data interface{}) error {
	// Load the template file.
	tmpl := template.Must(template.ParseFiles(templateFile))

	// Render the template file.
	return tmpl.Execute(w, data)
}

func signinHandler(fs http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The path `/` is a catch-all. To strictly match `/` to this handler,
		// we send other paths to the static handler.
		if r.URL.Path != "/" {
			fs.ServeHTTP(w, r)
			return
		}

		// GET  = display the sign-in form.
		// POST = handles the sign-in form submission.
		if r.Method == http.MethodGet {
			if err := renderTemplate(w, nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if r.Method == http.MethodPost {
			var msg strings.Builder

			// Get form values.
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			username := r.Form.Get("username")
			password := r.Form.Get("password")
			if username == "" {
				msg.WriteString("username is empty")
			}
			if password == "" {
				msg.WriteString("password is empty")
			}

			// Check credentials.
			if username != "sam" || password != "password" {
				msg.WriteString("invalid credentials")
			}

			if msg.Len() > 0 {
				content := &struct {
					Error string
				}{
					Error: msg.String(),
				}

				if err := renderTemplate(w, content); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				return
			}

			// Redirect if successful
			http.Redirect(w, r, successURI, http.StatusFound)
		} else {
			// Other methods
			http.Error(w, fmt.Sprintf("%s not supported", r.Method), http.StatusNotImplemented)
			return
		}
	})
}

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", signinHandler(fs))
	log.Println("web server running at port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
