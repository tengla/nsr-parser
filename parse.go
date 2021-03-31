package main

import (
	"encoding/xml"
	"os"
)

func StopPlaces(path string) (<-chan StopPlace, <-chan error) {
	ch := make(chan StopPlace)
	errCh := make(chan error)
	go func() {
		xmlFile, err := os.Open(path)
		if err != nil {
			errCh <- err
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
					decoder.DecodeElement(&sp, &el)
					ch <- sp
				}
			}
		}
		close(ch)
		close(errCh)
	}()
	return ch, errCh
}
