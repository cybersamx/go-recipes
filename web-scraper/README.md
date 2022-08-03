# Web Scrapping

**Note: There seems to be a bug in this example. The program is no longer scraping the content correctly. This will be fixed.**

This example shows how to scrap a website using the `gocolly/colly` package. Specifically, the code extracts a list of historical events that happened "on this day" from Wikipedia

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
