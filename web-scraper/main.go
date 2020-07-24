package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func parseSymbol() string {
	flag.Parse()

	if len(flag.Args()) == 0 {
		return ""
	}

	return strings.ToUpper(flag.Args()[0])
}

func extractPrice(e *colly.HTMLElement) {
	// Here's a sample of the DOM we want to extract the share price from:
	// <div class="intraday__close">
	//   <table class="table table--primary align--right">
	//     <thead>
	//       <tr class="table__row">
	//         <th class="table__heading">Close</th>
	//         <th class="table__heading">Chg</th>
	//         <th class="table__heading">Chg %</th>
	//       </tr>
	//     </thead>
	//     <tbody class="remove-last-border">
	//       <tr class="table__row">
	//         <td class="table__cell u-semi">$1,513.07</td>
	//         <td class="table__cell not-fixed negative">-79.26</td>
	//         <td class="table__cell not-fixed negative">-4.98%</td>
	//         <td class="table__cell fixed-to-top negative">-79.26 -4.98%</td>
	//       </tr>
	//     </tbody>
	//   </table>
	// </div>

	class := e.Attr("class")
	if strings.Contains(class, "intraday__close") {
		e.ForEach("table.table--primary tbody tr", func(_ int, tr *colly.HTMLElement) {
			fmt.Println("Close price:", tr.ChildText("td:first-child"))
			fmt.Println("Change:", tr.ChildText("td:nth-child(2)"))
			fmt.Println("Change%:", tr.ChildText("td:nth-child(3)"))
		})
	}
}

func main() {
	symbol := parseSymbol()
	if symbol == "" {
		fmt.Println("You must enter a ticker symbol")
		os.Exit(1)
	}

	fmt.Println("fetch stock quote for", symbol)

	c := colly.NewCollector(
		// Good practice to set the user-agent. See http://go-colly.org/articles/scraping_related_http_headers/
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.4 Safari/605.1.15"),
		// Visit only the following domains.
		colly.AllowedDomains("www.marketwatch.com", "marketwatch.com"),
		// Cache web pages.
		colly.CacheDir("./.cache"),
	)

	c.OnHTML("div", extractPrice)

	url := fmt.Sprintf("https://www.marketwatch.com/investing/stock/%s", symbol)
	if err := c.Visit(url); err != nil {
		log.Fatal(err)
	}
}
