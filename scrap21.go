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

	collector.OnHTML(".cocktail-item.promoted ", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("data-id"))
		factName := element.DOM.Find(".cocktail-item-preview").Text()
		factRecipe := element.DOM.Find(".cocktail-item-goods").Text()
		factEquipments := element.ChildAttr("a","class")
		if err != nil {
			log.Println("Could not get id")
		}

		//factDesk := element.DOM.Find("").Text()

		fact2 := Fact2{
			ID:         factId,
			Name:       factName,
			Recipes:    factRecipe,
			Equipments: factEquipments,
		}

		allFacts2 = append(allFacts2, fact2)

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

	writeJason2(allFacts2)
}

func writeJason2(data []Fact2) {
	file, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		log.Println("Unable create json file")
		return
	}

	_ = ioutil.WriteFile("jason_file2.json", file, 0644)
}
