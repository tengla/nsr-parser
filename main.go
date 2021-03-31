package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	stopPlaces, errs := StopPlaces("./nsr.current.xml")
	i := 0
	for stop := range stopPlaces {
		i += 1
		fmt.Println(stop.Name, stop.StopPlaceType, stop.TariffZones)
	}
	for err := range errs {
		log.Fatal(err)
	}
	duration := time.Since(start)
	fmt.Printf("StopPlace records: %d, duration: %v\n", i, duration)
}
