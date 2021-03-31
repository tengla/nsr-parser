package main

import (
	"flag"
	"fmt"
	"time"
)

var xmlfile = flag.String("xmlfile", "./nsr.current.xml", "Xml src file")

func main() {
	flag.Parse()
	start := time.Now()
	stops := StopPlaces(*xmlfile)
	i := 0
	for stop := range stops {
		i += 1
		switch t := stop.(type) {
		case error:
			fmt.Printf("Error: %s\n", t.Error())
		case StopPlace:
			fmt.Println(t.Name, t.StopPlaceType, t.TariffZones)
		}
	}
	duration := time.Since(start)
	fmt.Printf("StopPlace records: %d, duration: %v\n", i, duration)
}
