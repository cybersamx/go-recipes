# Go HTML Templating

This example shows server-side web content rendering using Go's `html/template` package.

Some highlights about `html/template` package:

* There's some context-aware safety. The package automatically escape all HTML template and inputs/content before they are rendered.
* You apply a `struct` that encapsulates the content to a template to produce the desired content output.

## Actions

Actions are control structures in the templates denoted by `{{` and `}}` and they are used to evaluate data or to control execution flow. Here are some of the actions

| Action                                             | Description                                                       |
|----------------------------------------------------|-------------------------------------------------------------------|
| `{{/* Comment */}}`                                | A comment                                                         |
| `{{pipeline}}`                                     | Renders the data representation. Can be a function, variable, etc |
| `{{.}}`                                            | Renders the entire data structure                                 |
| `{{.Field}}`                                       | Renders the field named Field of the data structure               |
| `{{functionName arg1 arg2 ...}}`                   | Call functions supported in Go Template                           |
| `{{.Method arg1 arg2 ...}}`                        | Call method with arguments                                        |
| `{{if pipeline}} {{else}} {{end}}`                 | Conditional statement on a boolean pipeline                       |
| `{{range $index, $item := pipeline}} {{}} {{end}}` | Loop over an array pipeline                                       |
| `{{range $key, $item := pipeline}} {{}} {{end}}`   | Loop over a map pipeline                                          |
| `{{with $variable := pipeline}} {{}} {{end}}`      | Define a variable                                                 |
| `{{define "name"}}template-content{{end}}`         | Define template content                                           |
| `{{template "name" pipeline}}`                     | Render the template named "name" with the passed pipeline         |

> **Functions**
>
> Some predefined functions in Go Template include:
> * `and`, `call`, `html`, `index`, `slice`, `js`, `len`, `not`, `or`, `print`, `printf`, `println`, `urlquery`.
> * There are also a set of functions that mimic binary comparison operators: `eq`, `ne`, `lt`, `le`, `gt`, `ge`.

## Setup

1. Run the server.

   ```bash
   $ make run
   ```

1. Launch a web browser and navigate to <http://localhost:8000>.

## Reference

* [Godoc: text/template](https://godoc.org/text/template)
* [Godoc: html/template](https://godoc.org/html/template)
