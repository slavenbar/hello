package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("ru.inshaker.com"),
	)
	c.OnHTML(".common-title header", func(e *colly.HTMLElement) {
		writer.Write([]string{
			//e.ChildText("h1"),
			e.ChildText("name"),
			//e.ChildAttr(""),
		})
	})

	for i := 0; i < 5; i++ {
		fmt.Printf("Scraping Page : %d\n", i)
		c.Visit("https://ru.inshaker.com/cocktails/" + strconv.Itoa(i))
	}

	log.Printf("Scraping Complete\n")
	log.Println(c)
}
