package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/saurabhmittal16/ticket/checker"
)

var url, target string

func init() {
	flag.StringVar(&url, "url", "https://in.bookmyshow.com/buytickets/joker-national-capital-region-ncr/movie-ncr-ET00100071-MT/20191002", "BookMyShow URL of movie")
	flag.StringVar(&target, "theatre", "INOX: Janak Place", "Name of theatre used in BookMyShow")
	flag.Parse()
}

func main() {
	ticker := time.NewTicker(5 * time.Minute)

	// for every tick in 5 minutes, check for availability
	for range ticker.C {
		found := checker.CheckAvailability(url, target)
		if found {
			fmt.Println("Tickets are available, book now " + url)
			os.Exit(0)
		} else {
			fmt.Println("Not available; checked at:", time.Now().Format("Jan 2 15:04:05 2006"))
		}
	}
}
