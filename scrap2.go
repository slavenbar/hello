package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"Description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("inshaker.com", "ru.inshaker.com"),
	)

	collector.OnHTML(".cocktail-item-preview", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("title"))
		if err != nil {
			log.Println("Could not get id")
		}

		factDesc := element.Text

		fact := Fact{
			ID:          factId,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)

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

    writeJason(allFacts)
}

func writeJason(data []Fact) {
	file, err := json.MarshalIndent(data,"", " ")
	if err != nil {
		log.Println("Unable create json file")
		return
	}
	
	_ = ioutil.WriteFile("jason_file.json", file, 0644)
}
