# Web Scrapping

A recipe that extracts a list of historical events that happened "on this day" from Wikipedia using the `gocolly/colly` package.

## Setup

1. Build the program.

   ```bash
   $ make
   ```

1. Run the program. You need to pass the program the --flush flag to flush out cache.

   ```bash
   $  bin/web-scraper --flush
   ```

## Reference

* [Colly](http://go-colly.org/)
* [Colly Doc](http://go-colly.org/docs)
