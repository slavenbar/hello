package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact2 struct {
	ID          int    `json:"id"`
	Description string `json:"Description"`
}

func main() {
	allFacts2 := make([]Fact2, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("inshaker.com", "ru.inshaker.com"),
	)

	collector.OnHTML(".cocktail-item.promoted", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("data-id"))
		if err != nil {
			log.Println("Could not get id")
		}

		factDesc := element.Text

		fact2 := Fact2{
			ID:          factId,
			Description: factDesc,
		}

		allFacts2 = append(allFacts2, fact2)

		for i := 0; i < 5; i++ {
			fmt.Printf("Scraping Page : %d\n", i)
			collector.Visit("https://ru.inshaker.com/cocktails/" + strconv.Itoa(i))
		}

		log.Printf("Scraping Complete\n")
		log.Println(collector)
	})
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})
	collector.Visit("https://ru.inshaker.com/cocktails/")

	writeJason2(allFacts2)
}

func writeJason2(data []Fact2) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable create json file")
		return
	}

	_ = ioutil.WriteFile("jason_file2.json", file, 0644)
}
