package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact3 struct {
	ID         int    `json:"id"`
	Name       string `json:"Name"`
	Recipes    string `json:"Recipes"`
	Equipments string `json:"Equipments"`
}

func main() {
	allFacts3 := make([]Fact3, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("inshaker.com", "ru.inshaker.com"),
	)

	collector.OnHTML(".ingredient-tables", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("data-id"))
		//factEquipments := element.DOM.Find("td:nth-child(2)").Text()
		factEquipments := element.ChildText("td")
		if err != nil {
			log.Println("Could not get id")
		}

		//factDesk := element.DOM.Find("").Text()

		fact3 := Fact3{
			ID:         factId,
			Equipments: factEquipments,
		}

		allFacts3 = append(allFacts3, fact3)

	})
	//".common-title header.common-name"

	for i := 0; i < 5; i++ {
		fmt.Printf("Scraping Page : %d\n", i)
		collector.Visit("https://ru.inshaker.com/cocktails/" + strconv.Itoa(i))
	}

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})
	collector.Visit("https://ru.inshaker.com/cocktails/")
	log.Printf("Scraping Complete\n")
	log.Println(collector)

	writeJason3(allFacts3)
}

func writeJason3(data []Fact3) {
	file, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		log.Println("Unable create json file")
		return
	}

	_ = ioutil.WriteFile("jason_file3.json", file, 0644)
}
