package checker

import (
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/gocolly/colly"
)

// CheckAvailability accepts the URL of BookMyShow and theatre name
// and checks the availability of tickets at that theatre
func CheckAvailability(url string, theatre string) {
	c := colly.NewCollector()

	c.OnHTML("a.__venue-name", func(e *colly.HTMLElement) {
		name := strings.TrimSpace(e.Text)
		if name == theatre {
			err := beeep.Notify("Tickets available", "Tickets are available, book now "+url, "assets/icon.jpg")
			if err != nil {
				panic(err)
			}
		}
	})

	c.Visit(url)
}
