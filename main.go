package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/saurabhmittal16/ticket/checker"
)

var url, target string
var duration int64

func init() {
	flag.StringVar(&url, "url", "https://in.bookmyshow.com/buytickets/joker-national-capital-region-ncr/movie-ncr-ET00100071-MT/20191002", "BookMyShow URL of movie")
	flag.StringVar(&target, "theatre", "INOX: Janak Place", "Name of theatre used in BookMyShow")
	flag.Int64Var(&duration, "duration", 30, "Duration after which tickets are checked")
	flag.Parse()
}

func main() {
	ticker := time.NewTicker(time.Duration(duration) * time.Minute)
	fmt.Printf("Checking for %v at %v in %v minute(s)\n", url, target, duration)

	// for every tick in given minutes, check for availability
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
