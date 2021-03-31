package main

import (
	"encoding/xml"
	"os"
)

func StopPlaces(path string) chan interface{} {
	ch := make(chan interface{})
	go func() {
		xmlFile, err := os.Open(path)
		if err != nil {
			ch <- err
			close(ch)
			return
		}
		decoder := xml.NewDecoder(xmlFile)
		for {
			t, _ := decoder.Token()
			if t == nil {
				break
			}
			switch el := t.(type) {
			case xml.StartElement:
				if el.Name.Local == "StopPlace" {
					var sp StopPlace
					if err = decoder.DecodeElement(&sp, &el); err != nil {
						ch <- err
					} else {
						ch <- sp
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}
