# Go Module

## Enable Vendor Mode

If you want to maintain a copy of all third party packages locally in your project - similar to `node_modules` in JavaScript, you need to enable vendoring mode in your Go module. To do that, you run the following command:

```bash
$ # Initialize a new module
$ go init github.com/cybersamx/go-app
$ # Enable vendor
$ go mod vendor
```

This will create a vendor directory in the project. I would add the directory `vendor` to the project's `.gitignore`.

## Access Local Package

If you have a source code for a package that exists locally, you may configure your go.mod the following way to reference the local copy instead of pulling the package from the Internet.

```go
// file: /pkg/go.mod
module example.com/me/app

require (
  example.com/me/pkg v0.0.0
)

replace example.com/me/pkg => ../pkg
```

This assumes the following directory structure:

```
- pkg
   - go.mod
- app
   - go.mod
```

## 410 Gone Issue

Sometimes you may reference a package that is hosted in a private repo. You may get an `410 Gone` error as you download the package when you building your project. To resolve this, you need to indicate that you are using a private repo by doing the following:

```bash
$ go mod tidy
...
verifying github.com/private/lib... reading https://sum.golang.org/lookup/github.com/private/lib@: 410 Gone
$ # Do the following to resolve the issue
$ export GOPRIVATE=github.com/private/lib
$ go mod tidy
```
