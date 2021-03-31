package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"
)

var (
	xmlfile = flag.String("xmlfile", "./nsr.current.xml", "Xml src file")
	jsonout = flag.Bool("jsonout", false, "output json")
)

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
			if *jsonout {
				j, _ := json.Marshal(t)
				fmt.Println(string(j))
			} else {
				fmt.Printf("Id=%s, Name=%s, StopPlaceType=%s,",
					t.Id, t.Name, t.StopPlaceType)
				for _, list := range t.KeyList {
					for _, kv := range list.KeyValue {
						fmt.Printf("%s=%s, ", kv.Key, kv.Value)
					}
				}
				fmt.Println()
			}
		}
	}
	duration := time.Since(start)
	if !*jsonout {
		fmt.Printf("\nStopPlace records: %d, duration: %v\n", i, duration)
	}
}
