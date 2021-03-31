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
				fmt.Printf("Id=%s,Name=%s,StopPlaceType=%s\n",
					t.Id, t.Name, t.StopPlaceType)
			}
		}
	}
	duration := time.Since(start)
	if !*jsonout {
		fmt.Printf("StopPlace records: %d, duration: %v\n", i, duration)
	}
}
