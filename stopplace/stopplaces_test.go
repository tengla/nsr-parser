package stopplace

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func TestStopPlaces(t *testing.T) {
	t.Run("StopPlaces", func(t *testing.T) {
		d, _ := os.Getwd()
		file := d + "/../nsr.test.xml"
		ctx, cancel := context.WithTimeout(
			context.Background(),
			// dont exceed 30 seconds
			time.Duration(time.Second*30))
		defer cancel()
		stops := StopPlaces(ctx, file)
		found := false
		for stop := range stops {
			switch t := stop.(type) {
			case StopPlace:
				if t.ID == "NSR:StopPlace:5" {
					found = true
				}
			}
		}
		if !found {
			log.Fatal("No records found")
		}
	})
}
