# HTTP Cookies in Go

A recipe on working with http cookies in Go. Cookie is a simple construct but its usage has huge implications as understanding what a cookie does and how to implement it properly in Go is critical to developing secure web applications.

## Cookies Overview

A cookie is nothing more than a key-value pair that is stored on a client web browser. Typically a server sends a request using the HTTP `Set-Cookie` directive in the response header to request that a key-value pair be saved on the client. Subsequent requests from the web browser to the same server will include the same cookie in the `Cookie` HTTP header. 

This recipe create several different cookies to illustrate some key concepts:

### HttpOnly Field
 
Cookie with `HttpOnly` field set to `true` makes the cookie inaccessible to the Javascript `document.cookie` on the client. The cookie will still be sent back to the server in the `Cookie` HTTP header, hence the name `HttpOnly`. This is useful if you want to persist strictly server-side data that don't need to be exposed to the client Javascript eg. server-side sessions.

### Secure Field

When the `Secure` field is set to `true`, we tell the web browser that the cookie must be sent to the server over HTTPS.

### Domain Field

When a cookie is set to a specific domain, the client web browser will only send the cookie if the URL matches the specified domain. If the domain attribute isn't set, the domain defaults to the host of the current document location (but excluding its subdomain). However if the domain is set, the subdomains are always included. For example, if we set `Domain = github.com`, the cookie is available to the web browser if the URL is `github.com` or `sub.github.com`. See [here](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies) for details.

### Path Field

Another way to control the scope of the cookie is path. When a cookie is set to a specific path, the client web browser will only send the cookie if the URL includes the specified path. All subdirectories of the specified path will be included as well. For example, setting `Path = /images` means that the cookie will be sent for the following paths if found in the URL:

* `/images`
* `/images/a`
* `/images/a/acme`

See [here](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies) for more details.

### SameSite Field

There is now a new field called `SameStie` that can be used to control the scope of where a cookie should be sent. The field accepts 3 enum values: None, Strict, Lax.

> **Notes**
>
> * Not all web browser supports `SameSite` - [See here for browser support](https://developer.mozilla.org/en-US/docs/Web/HTTP/headers/Set-Cookie#Browser_compatibility).
> * The `SameSite` field is only available in Go 1.11 and above.

## MaxAge vs Expires Fields

[This article explains this subject perfectly](https://mrcoles.com/blog/cookies-max-age-vs-expires/). The `Expires` and `MaxAge` attributes denotes the same thing, that is when a cookie will expire. The former uses an absolute timestamp (when the cookie will expires at a specific timestamp) and the latter a relative duration (when the cookie will expires after a specified duration). `Expires` was deprecated back in 2009. So just use `MaxAge`. 

> **Notes**
>
> * Setting `MaxAge < 0` tells the web browser to delete the cookie.
> * MaxAge is in number of seconds.

### Session vs Permanent Cookies

A session cookie is defined as a cookie that is created without specifying an `Expires` or `MaxAge` field. When a user quits a web browser, all session cookies are deleted. A permanent cookie is one that has either an `Expires` or `MaxAge` field set.

### Serialization and Encoding

The `Value` attribute of a cookie accepts only string value. Nonetheless you can serialized any Golang types into a string and then encode it as Base64 before saving it to the cookie.

### Security

The [securecookie package](https://github.com/gorilla/securecookie), as part of the Gorilla Toolkit for writing robust, secure web apps in Go, has a good API for writing secure cookie.

## Setup

1. Run the program

   ```bash
   $ go run main.go
   ```
   
1. Open a web browser and navigate to <http://localhost:8000/vend> for the application to vend out cookies.

1. Open the developer tool/console and see what cookies are set. You can also go to <http://localhost:8000/read> to have the server print out all the cookies.

   ![Chrome](images/chrome-developer-tool.png)
   
## Reference

* [Morzilla: Cookie](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies)
* [Godoc: Cookie struct](https://golang.org/pkg/net/http/#Cookie)
* [Gorilla Web Toolkit: securecookie](https://www.gorillatoolkit.org/pkg/securecookie)
