package main

import (
	"fmt"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/gocolly/colly"
)

func main() {
	// url := "https://in.bookmyshow.com/buytickets/joker-national-capital-region-ncr/movie-ncr-ET00100071-MT/20191002"
	url := "https://in.bookmyshow.com/buytickets/dream-girl-national-capital-region-ncr/movie-ncr-ET00089745-MT/20190930"
	target := "INOX: Janak Place"

	c := colly.NewCollector()

	c.OnHTML("a.__venue-name", func(e *colly.HTMLElement) {
		name := strings.TrimSpace(e.Text)
		if name == target {
			fmt.Println("Found")
			err := beeep.Notify("Tickets available", "Tickets are available, book now "+url, "assets/icon.jpg")
			if err != nil {
				panic(err)
			}
		}
	})

	c.Visit(url)
}
