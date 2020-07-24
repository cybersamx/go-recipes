# Web Scrapping

A recipe that extracts the price quote of a stock from MarketWatch using the `gocolly/colly` package.

## Setup

1. Build the program.

   ```bash
   $ make
   ```

1. Run the program. You need to pass the program a ticker symbol eg. make run [ticker symbol].

   ```bash
   $  bin/web-scraper GOOG
   ```

## Reference

* [Colly](http://go-colly.org/)
* [Colly Doc](http://go-colly.org/docs)
