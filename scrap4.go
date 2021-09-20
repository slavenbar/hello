package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "inshaker.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Receipes", "Description"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("#currencies-all tbody tr", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("ingredient-tables"),
		})
	})

	c.Visit("https://ru.inshaker.com/cocktails/31-krovavaya-meri")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
