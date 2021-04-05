package stopplace

import (
	"context"
	"encoding/xml"
	"errors"
	"os"
	"time"
)

func StopPlaces(ctx context.Context, path string) chan interface{} {
	ch := make(chan interface{})
	deadline, ok := ctx.Deadline()
	go func() {
		if !ok {
			ch <- errors.New("context needs to be set with a deadline")
			close(ch)
			return
		}
		xmlFile, err := os.Open(path)
		if err != nil {
			ch <- err
			close(ch)
			return
		}
		decoder := xml.NewDecoder(xmlFile)
		for {
			if time.Now().After(deadline) {
				break
			}
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
