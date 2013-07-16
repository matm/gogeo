package main

import (
	"flag"
	"fmt"
	"github.com/matm/gogeo"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [option] full_street_address\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	mqprov := flag.Bool("mq", false, "use MapQuest geocoding API")
	gprov := flag.Bool("g", false, "use Google geocoding API")

	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprint(os.Stderr, "Only one argument allowed for street address.\n")
		flag.Usage()
	}

	var gc gogeo.GeoCoder

	if *gprov {
		gc = gogeo.NewGoogleGeoCoder()
	} else if *mqprov {
		gc = gogeo.NewMapQuestGeoCoder()
	} else {
		fmt.Fprintf(os.Stderr, "Please specify a request provider.\n")
		flag.Usage()
	}

	gps, err := gc.Geocode(&gogeo.Location{FullAddress: flag.Arg(0)})
	if err != nil {
		fmt.Printf("geocoding: %s", err.Error())
	}
	if gps == nil {
		fmt.Println("no GPS coordinates returned")
		os.Exit(3)
	}
	fmt.Println(gps)
}
