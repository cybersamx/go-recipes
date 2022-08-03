# Layouts (Multi-Templates and Parameters)

As we scale our website, we will likely be embedding a html template into another template so that we can reuse code and content.

This example demonstrates 2 things:

1. Parse and load multiple templates.
1. Pass data to an embedded template.

## Loading Multiple Templates

Multiple templates can be parsed in one of the following ways:

* Use both `filepath.Glob` and `template.ParseFiles`
* Use `template.ParseGlob`, which is a convenience function of calling the above 2 functions.

Pass the return value from the above function call to `template.Must`, which will return a `template.Template` value.

We use the `ExecuteTemplate` method in the `template.Template` value to render a specific named template. To make this work, we must define a template name in the template file `*.gohtml` with the directive `{{define "myTemplateName"}}`.

## Passing Data to an Embedded Template

If you look at the code, you see that we embed a template named `content` into the `home` template by using this directive `{{template "content" .Data}}`. This directive tells Go runtime to render the content from the template `content` along with the passed parametric data `.Data` onto the parent template `home`. The problem is that we can only pass 0 or 1 parameter to the template. So if you need to pass more than 1 value, you would need to wrap all the multiple values as a complex data type.

## Setup

1. Run the server.

   ```bash
   $ make run
   ```

1. Launch a web browser and navigate to <http://localhost:8000>.

## Reference

* [Godoc: text/template](https://godoc.org/text/template)
* [Godoc: html/template](https://godoc.org/html/template)
