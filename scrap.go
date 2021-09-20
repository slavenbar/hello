package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	scrapPage("https://ru.inshaker.com/cocktails/35-long-aylend-ays-ti")
}

func scrapPage(url string) {
	c := colly.NewCollector()

	c.OnHTML("td > table", func(e *colly.HTMLElement) {
		enw := e.DOM.Find("td:nth-child(2)").Text()
		fmt.Println(enw)
	})
	c.Visit(url)
}
