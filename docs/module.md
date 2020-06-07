# Go Module

## Initialize a New Module

Initialize your project as a module. Basically creates a `go.mod` in your project file. If you don't pass a module path as an argument, the program will guess the path for you based on the version control config and directory structure.

```bash
$ go mod init
```

If you have a module path in mind, just pass it as an argument.

```bash
$ go mod init github.com/cybersamx/go-recipes
```

## Tidy Up Your Module

To make sure that your `go.mod` matches the modules to be imported in your source code, run the following command to add missing or remove unneeded modules.

```bash
$ go mod tidy -v
```

## Enable Vendor Mode

If you want to maintain a copy of all third party modules (dependencies) locally in your project - similar to `node_modules` in JavaScript, you need to enable vendor mode in your Go module. To do that, you run the following command:

```bash
$ # Initialize a new module
$ go init github.com/cybersamx/go-app
$ # Enable vendor
$ go mod vendor
```

This will create a vendor directory in the project. I would add the directory `vendor` to the project's `.gitignore`.

## List Modules

List all third-party modules (or dependencies) required by a module.

```bash
$ go list -m all
$ go mod graph
```

If you wonder why a particular module shows up in your go.mod file, run this command:

```bash
$ go mod why -m <module>
```

## Use Local Module

If you have a source code for a module that exists locally, you may configure your go.mod the following way to reference the local copy instead of pulling the module from the Internet.

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
   - go.mod   <- example.com/me/app
- app
   - go.mod   <- example.com/me/pkg
```

> **Notes**
>
> The directive `replace` in `go.mod` instructs go to replace a module version with a different module version. We are using to override the default behavior to use the local copy instead of the remote copy.

## 410 Gone Issue

Sometimes you may reference a module that is hosted in a private repo. You may get an `410 Gone` error as you download the package when you building your project. To resolve this, you need to indicate that you are using a private repo by doing the following:

```bash
$ go mod tidy
...
verifying github.com/private/lib... reading https://sum.golang.org/lookup/github.com/private/lib@: 410 Gone
$ # Do the following to resolve the issue
$ export GOPRIVATE=github.com/private/lib
$ go mod tidy
```
