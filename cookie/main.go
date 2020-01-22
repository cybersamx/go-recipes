package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	addr = ":8000"
	key  = "abcdefghijklmnopqrstuvwxyz123456"
	expire = 60 * time.Minute
)

func getSimpleCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:       name,
		Value:      value,
		Expires:    time.Now().Add(expire),
	}
}

func getHTTPOnlyCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:       name,
		Value:      value,
		Expires:    time.Now().Add(expire),
		HttpOnly:   true,
	}
}

func getDomainCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:       name,
		Value:      value,
		Expires:    time.Now().Add(expire),
		Domain:     "example.com",
	}
}

func getPathCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:       name,
		Value:      value,
		Path:       "/nonexistent",
		Expires:    time.Now().Add(expire),
	}
}

type entry struct{
	name  string
	value string
	getter func(name, value string) *http.Cookie
}

func vendCookieHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// --- Create some cookies ---
		entries := []entry{
			{name: "simple", value: "simple value", getter: getSimpleCookie},       // 1. Get simple cookie.
			{name: "http-only", value: "simple value", getter: getHTTPOnlyCookie},  // 2. Get http-only cookie.
			{name: "domain", value: "domain value", getter: getDomainCookie},       // 3. Get domain-specific cookie.
			{name: "path", value: "path value", getter: getPathCookie},             // 4. Get path-specific cookie.
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
		w.WriteHeader(http.StatusOK)

		for _, cookie := range r.Cookies() {
			if _, err := io.WriteString(w, cookie.Value + "\n"); err != nil {
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
