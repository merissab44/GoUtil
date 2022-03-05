package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

// create shoe struct
type shoe struct {
	Name  string
	Price string
}

// create json function
func createJson(shoeList []shoe) {
	jsonFile, _ := json.MarshalIndent(shoeList, "", " ")
	_ = ioutil.WriteFile("shoes.json", jsonFile, 0644)
}

// create collection
func main() {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	shoes := make([]shoe, 0)

	// On every a element which has href attribute call callback
	c.OnHTML("div.css-xzkzsa", func(e *colly.HTMLElement) {
		e.ForEach("div.css-1ibvugw-GridProductTileContainer", func(_ int, elm *colly.HTMLElement) {
			newShoe := shoe{}
			newShoe.Name = elm.ChildText("p.css-3lpefb")
			newShoe.Price = elm.ChildText("p.css-9ryi0c")
			shoes = append(shoes, newShoe)
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got Error:", e)
	})

	// Start scraping on stockc.com
	c.Visit("https://stockx.com/sneakers/most-popular")

	createJson(shoes)

	// ask user to input their budget
	var budget float64
	fmt.Println("Please enter your budget:")
	fmt.Scanf("%f", &budget)

	for _, shoe := range shoes {
		price, _ := strconv.ParseFloat(shoe.Price[1:], 64)
		if price < budget {
			fmt.Println("These are the shoes are within your budget", shoe.Name, shoe.Price)
		}
	}
}
