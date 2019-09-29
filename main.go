package main

import (
	"flag"
	"stark/ticket/checker"
)

var url, target string

func init() {
	flag.StringVar(&url, "url", "https://in.bookmyshow.com/buytickets/joker-national-capital-region-ncr/movie-ncr-ET00100071-MT/20191002", "BookMyShow URL of movie")
	flag.StringVar(&target, "theatre", "INOX: Janak Place", "Name of theatre used in BookMyShow")
	flag.Parse()
}

func main() {
	checker.CheckAvailability(url, target)
}

// go run main.go --url="https://in.bookmyshow.com/buytickets/dream-girl-national-capital-region-ncr/movie-ncr-ET00089745-MT/20190930" --theatre="INOX: Janak Place"
