# HTTP Recipes

Recipes that shows HTTP client and server implementations.

## Client

* [Client](client) - a recipe that implements a simple HTTP client that sends a GET request.

## Server

* [JWT](server/jwt) - a recipe that shows the handling of simple web form submission and a simple authentication using JWT.
* [Static HTTP server](server/static) - a recipe that implements a simple http server serving static content.

### Notes

#### Catch-All Routing

A path "pattern ending in a slash names a rooted subtree." This means that a root path ie. `/` matches all paths that is not handled by any handler. In a way think of `/` as a catch-all for all paths. See: [GoDocs: ServerMux](https://pkg.go.dev/net/http#ServeMux) for details.

```go
// If you have this in your code and there are no other handlers.
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<html><body>Hello World</body></html>")
})
```

Calling the following urls will return the same "Hello World" response.

```bash
$ curl http://localhost:8000/
$ curl http://localhost:8000
$ curl http://localhost:8000/about
$ curl http://localhost:8000/somethingThatDoesNotExist
```

And if you really want only `/` to return "Hello World", you now have a problem. The workaround is to declare and register the paths with more handlers or just use this pattern.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    io.WriteString(w, "<html><body>Hello World</body></html>")
})
```
