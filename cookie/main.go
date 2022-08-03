package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

const (
	addr   = ":8000"
	expire = 24 * 60 * 60
)

var (
	hashKey  = []byte("abcdefghijklmnopqrstuvwxyz123456")
	blockKey = []byte("kcaur4fi3r7DSFasjfwer3274we9r737")
)

func getCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: expire,
	}
}

func getHTTPOnlyCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   expire,
		HttpOnly: true,
	}
}

func getDomainCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: expire,
		Domain: "example.com",
	}
}

func getPathCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:   name,
		Value:  value,
		Path:   "/nonexistent",
		MaxAge: expire,
	}
}

type entry struct {
	name   string
	value  string
	getter func(name, value string) *http.Cookie
}

type Preferences struct {
	Font     string
	TextSize int
	BGColor  string
	Margin   int
}

func serializePreferences(preferences *Preferences) (string, error) {
	gob.Register(&Preferences{})
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(&preferences); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func vendCookieHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Encode a complex type to string.
		serialized, err := serializePreferences(&Preferences{
			Font:     "Courier",
			TextSize: 16,
			BGColor:  "Silver",
			Margin:   4,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Hash a cookie.
		sc := securecookie.New(hashKey, blockKey)
		hashed, err := sc.Encode("hash", "hash-value")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// --- Create some cookies ---
		entries := []entry{
			{name: "standard", value: "standard-value", getter: getCookie},      // Simple cookie.
			{name: "http-only", value: "http-value", getter: getHTTPOnlyCookie}, // HTTP-only cookie.
			{name: "domain", value: "domain-value", getter: getDomainCookie},    // Domain-specific cookie.
			{name: "path", value: "path-value", getter: getPathCookie},          // Path-specific cookie.
			{name: "encoded", value: serialized, getter: getCookie},             // Standard cookie with serialization.
			{name: "hash", value: hashed, getter: getCookie},                    // Standard cookie with hashing.
		}

		for _, entry := range entries {
			http.SetCookie(w, entry.getter(entry.name, entry.value))
		}

		// Make sure that cookie is set before writing to the header and body. Otherwise
		// the cookie won't be set properly.
		w.WriteHeader(http.StatusOK)
	})
}

func readCookieHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, cookie := range r.Cookies() {
			if _, err := io.WriteString(w, cookie.Value+"\n"); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/vend", vendCookieHandler())
	mux.Handle("/read", readCookieHandler())

	log.Printf("web server running at %s", addr)
	log.Fatalf("web server failed: %v", http.ListenAndServe(addr, mux))
}
