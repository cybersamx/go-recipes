# GraphQL Server in Go

A recipe that implements a simple GraphQL server using Go and the [99 Designs GraphQL](https://gqlgen.com/) package and code generator.

## Setup

Here are the steps describing how we would build this project from scratch using `gqlqen`. For details, please read [gqlgen tutorial](https://gqlgen.com/getting-started/). Here's a summary of the tutorial:

1. Download the codegen tool.

   ```bash
   $ go get github.com/99designs/gqlgen
   ```

1. Create and initialize the project directory.

   ```bash
   $ mkdir server
   $ cd server
   $ go mod init github.com/cybersamx/go-recipes/graphql/server
   ```

1. Build the skeletal code using `gqlgen`. Note: this should generate some boilerplate, and a graphql server with a playground UI for experimenting with  graphql queries.

   ```bash
   $ gqlgen init   # Create skeletal code
   ```

1. Run the server.

   ```bash
   $ go run server.go
   ```

   Launch a web browser and go to http://localhost:8080. You should see the following page. Run the query on the browser:

   ```
   query {
     todos {
       id
     }
   }
   ```

   ![gqlgen playground](images/gqlgen-playground.png)

   Since we haven't implemented code to return the dataset, we get an error from the server.

1. Implement code to track state of the TODOs. Change the following files:

   * [resolver.go](graph/resolver.go)
   * [schema.resolvers.go](graph/schema.resolvers.go)

   See [here](https://gqlgen.com/getting-started/) for details.

1. Restart the server and rerun the previous graphql query. You should now get this response:

   ```
   {
     "data": {
       "todos": []
     }
   }
   ```

   Mutate the state (ie. add a TODO) to the server using this:

   ```
   mutation {
     createTodo(input: {
       text: "hello world",
       userId: "abc-123"
     }) {
       id,
       text
     }
   }
   ```

1. Add the following line to [resolver.go](graph/resolver.go).

   ```go
   //go:generate go run github.com/99designs/gqlgen
   ```

   This line instructs the go compiler what to do when we run it with the generate command recursively over the directory.

   ```bash
   $ generate ./...
   ```

All of this is captured in the [Makefile](Makefile). So simply run this command to build and run the project.

```bash
$ make
```

# Reference

* [gqlgen](https://gqlgen.com/)
