# Multi-Templates

In large production websites, we embed reusable templates (aka layouts) to make the website more scalable. For example, we would have 1 layout for the header and 1 layout for the footer; and we embed these templates into the main template.

This recipe shows how we embed `header.gohtml` and `footer.gohtml` layouts into the main template `home.gohtml`.

## Setup

1. Run the server

   ```bash
   $ make run
   ```

1. Launch a web browser and navigate to <http://localhost:8000>.

## Reference

* [Godoc: text/template](https://godoc.org/text/template)
* [Godoc: html/template](https://godoc.org/html/template)
