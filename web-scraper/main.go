package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

const (
	cacheDirPath = "./.cache"
)

func parseFlag() bool {
	flush := flag.Bool("flush", false, "flush cache")
	flag.Parse()

	return *flush
}

func removeCache(dirPath string) error {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) || !info.IsDir() {
		return nil
	}

	return os.RemoveAll(dirPath)
}

func today() string {
	now := time.Now()
	_, month, day := now.Date()
	return fmt.Sprintf("%s_%d", month, day)
}

func main() {
	if parseFlag() {
		fmt.Println("Flushing", cacheDirPath)
		if err := removeCache(cacheDirPath); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Extracting events for on this day from ")

	c := *colly.NewCollector(
		// Good practice to set the user-agent. See http://go-colly.org/articles/scraping_related_http_headers/
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.4 Safari/605.1.15"),
		// Visit only the following domains.
		colly.AllowedDomains("wikipedia.org", "en.wikipedia.org"),
		// Cache web pages.
		colly.CacheDir(cacheDirPath),
	)

	c.OnRequest(func(req *colly.Request) {
		// Set the request to the server.
	})

	c.OnResponse(func(res *colly.Response) {
		// Get and examine the response from the server.
	})

	c.OnHTML("div.mw-parser-output", func(e *colly.HTMLElement) {
		// Under div.mw-parser-output for for a h2 that is entitled Events.
		// Once found, extract the listing from the sibling ul element.
		ch := e.DOM.Children()
		for i := 0; i < ch.Size(); i++ {
			if ch.Get(i).Data == "h2" && ch.Eq(i).Text() == "Events" {
				// The next sibling has all the events.
				if i < ch.Size() - 1 {
					ul := ch.Eq(i + 1)
					for j := 0; j < ul.Children().Size(); j++ {
						li := ul.Children().Eq(j)
						fmt.Println(li.Text())
					}
				}
			}
		}
	})

	c.OnError(func(res *colly.Response, err error) {
		fmt.Errorf("error: %v", err)
	})

	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", today())
	if err := c.Visit(url); err != nil {
		log.Fatal(err)
	}
}
