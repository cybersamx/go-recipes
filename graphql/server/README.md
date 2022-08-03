# GraphQL Server in Go

An example of implementomg a simple GraphQL server using the [99 Designs GraphQL](https://gqlgen.com/) package.

## Setup

Here are the steps describing how we would build this project from scratch using `gqlqen`. For details, please read [gqlgen tutorial](https://gqlgen.com/getting-started/). Here's a summary of the tutorial:

Here's a sequence of steps to run the example:

1. Download the codegen tool.

   ```bash
   $ go install github.com/99designs/gqlgen@latest
   ```

1. Build the skeletal code using `gqlgen`. Note: this should generate some boilerplate, and a graphql server with a playground UI for experimenting with  graphql queries.

   ```bash
   $ cd <project_root_dir>
   $ gqlgen init   # Create skeletal code
   ```

Or we can build everything up to this point by running the following command:

```bash
$ make      # Build the project
```

1. Insert code to the skeletal code. In `graph/resolver.go`:

   ```go
   import (
     "github.com/cybersamx/go-recipes/graphql/server/graph/model"
   )

   // This file will not be regenerated automatically.
   //
   // It serves as dependency injection for your app, add any dependencies you require here.

   type Resolver struct {
     todos []*model.Todo
   }
   ```

   In `graph/schema.resolvers.go`:

   ```go
   import (
     "context"
     "fmt"
     "math/rand"

     "github.com/cybersamx/go-recipes/graphql/server/graph/generated"
     "github.com/cybersamx/go-recipes/graphql/server/graph/model"
   )

   // CreateTodo is the resolver for the createTodo field.
   func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
     todo := &model.Todo{
       Text: input.Text,
       ID:   fmt.Sprintf("T%d", rand.Int()),
       User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
     }
     r.todos = append(r.todos, todo)
     return todo, nil
   }

   // Todos is the resolver for the todos field.
   func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
     return r.todos, nil
   }
   ```

1. Run the server.

   ```bash
   $ go run server.go
   ```

   Launch a web browser and go to <http://localhost:8080>. You should see the following page. Copy and paste the following query on the browser:

   ```json
   query {
     todos {
       id
     }
   }
   ```

   ![gqlgen playground](images/gqlgen-playground.png)

   We should now get this response:

   ```json
   {
     "data": {
       "todos": []
     }
   }
   ```

   It's empty, so mutate the state (ie. add a TODO) to the server using this. Create another tab, copy and paste the following content, and run the graphql.

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
   $ go generate ./...
   ```

Once all the skeletal code is properly generated and set up, we can subsequently run the following to run the program.

```go
$ make run
```

# Reference

* [gqlgen](https://gqlgen.com/)
