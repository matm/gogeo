package main

import (
	"fmt"
	"github.com/matm/gogeo"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s full_address\n", os.Args[0])
		os.Exit(2)
	}

	gc := gogeo.NewMapQuestGeoCoder()
	gps, err := gc.Geocode(&gogeo.Location{FullAddress: os.Args[1]})
	if err != nil {
		fmt.Printf("geocoding: %s", err.Error())
	}
	if gps == nil {
		fmt.Println("no GPS coordinates returned")
		os.Exit(3)
	}
	fmt.Println(gps)
}
