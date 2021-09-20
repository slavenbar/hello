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
	ID         int    `json:"id"`
	Name       string `json:"Name"`
	Recipes    string `json:"Recipes"`
	Equipments string `json:"Equipments"`
}

func main() {
	allFacts2 := make([]Fact2, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("inshaker.com", "ru.inshaker.com"),
	)

	collector.OnHTML("td > table tr", func(element *colly.HTMLElement) {
		factName := element.DOM.Find("td:nth-child(2)").Text()

		factDesk := element.Text

		fact2 := Fact2{
			ID:   factId,
			Name: factName,
		}

		allFacts2 = append(allFacts2, fact2)

		for i := 0; i < 5; i++ {
			fmt.Printf("Scraping Page : %d\n", i)
			collector.Visit("https://ru.inshaker.com/cocktails/" + strconv.Itoa(i))
		}
	})
	//".common-title header.common-name"

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})
	collector.Visit("https://ru.inshaker.com/cocktails/")
	log.Printf("Scraping Complete\n")
	log.Println(collector)

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
