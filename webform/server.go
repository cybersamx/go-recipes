package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"
)

// --- Global Variables ---

type JWTClaims struct {
	Username string `json: "username"`
	jwt.StandardClaims
}

var (
	port = "8000"
	userContextKey = "username"
	expireDuration = 3 * time.Hour
	cookieName     = "jwt"
	secret         = []byte("Secret")
	users          = map[string]string{
		"user1": "password",
		"user2": "password",
	}
)

// --- Helper Functions ---

// Authenticate given username and password. Return true if user is authenticated successfully,
// otherwise returns an false with an appropriate error.

func authenticate(username string, password string) (bool, error) {
	foundPwd := users[username]
	if foundPwd == "" {
		return false, errors.New(fmt.Sprintf("user %s not found", username))
	}

	if foundPwd != password {
		return false, errors.New("password does not match")
	}

	return true, nil
}

// Read from an HTML template from the file system and stream the file's content to the
// HTTP response.

func outputHTML(w http.ResponseWriter, r *http.Request, filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	http.ServeContent(w, r, file.Name(), time.Now(), file)
}

func newJWT(username string) (string, error) {
	claims := JWTClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// JWT time format is in epoch.
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(secret)
}

// --- HTTP Filters ---

// HTTP filter to parse and extract the user name from the request. Only allow
// authenticated users to access this resource.

func restrictAccessFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get cookie
		cookie, err := r.Cookie(cookieName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			next.ServeHTTP(w, r)
			return
		}

		tokenStr := cookie.Value

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			next.ServeHTTP(w, r)
			return
		}

		// Validate and extract the username from the token.
		claims := JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			next.ServeHTTP(w, r)
			return
		}
		if !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			next.ServeHTTP(w, r)
			return
		}

		username := claims.Username

		// Add user to the context
		ctx := context.WithValue(r.Context(), userContextKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// --- HTTP Handlers ---

// Handles the login (web) form.

func loginHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")
			ok, err := authenticate(username, password)
			if !ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Generate JWT and  set in the response's cookie.
			tokenStr, err := newJWT(username)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:       cookieName,
				Value:      tokenStr,
				Expires:    time.Now().Add(expireDuration),
			})

			http.Redirect(w, r, "/welcome", http.StatusSeeOther)

			return
		}

		outputHTML(w, r, "html/login.html")
	})
}

func logoutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:    cookieName,
			Value:   "",
			Expires: time.Now().Add(-expireDuration),
		})

		outputHTML(w, r, "html/logout.html")
	})
}

func welcomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var noCacheHeaders = map[string]string{
			"Cache-Control": "no-cache, no-store, must-revalidate, max-age=0",
			"Pragma":        "no-cache",
			"Expires":       time.Unix(0, 0).Format(time.RFC1123),
		}

		// Set caching headers
		for k, v := range noCacheHeaders {
			w.Header().Set(k, v)
		}

		// Remove Etag header
		r.Header.Del("Etag")

		username := r.Context().Value(userContextKey)
		if username != nil {
			log.Printf("user %s accessing welcome page", username)
			outputHTML(w, r, "html/welcome.html")
		}
	})
}

// --- Main Function ---

func main() {
	http.Handle("/", loginHandler())
	http.Handle("/welcome", restrictAccessFilter(welcomeHandler()))
	http.Handle("/logout", logoutHandler())

	log.Println("Web server running at port", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

