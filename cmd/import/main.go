package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/tengla/nsr-parser/stopplace"
)

var (
	xmlfile = flag.String("file", "./nsr.current.xml", "XML source file")
	jsonout = flag.Bool("jsonout", false, "JSON output")
)

// importRecords
func importRecords(file string, jsonout bool) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		// dont exceed 30 seconds
		time.Duration(time.Second*30))
	defer cancel()

	start := time.Now()
	stops := stopplace.StopPlaces(ctx, file)
	count := 0

	for stop := range stops {
		count += 1
		switch t := stop.(type) {
		case error:
			fmt.Printf("Error: %s\n", t.Error())
		case stopplace.StopPlace:
			if jsonout {
				j, _ := json.Marshal(t)
				fmt.Println(string(j))
			} else {
				fmt.Printf("StopPlace(ID: %s, %s)\n", t.ID, t.Name.Text)
			}
		}
	}
	duration := time.Since(start)
	if !jsonout {
		fmt.Printf("\nRecord count: %d, duration: %s\n", count, duration)
	}
}
func b2mb(n uint64) float64 {
	return float64(n) / 1024 / 1024
}

func printMemUsage() string {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	s := fmt.Sprintf("Alloc = %.2f Mb", b2mb(m.Alloc))
	s = s + fmt.Sprintf("\tTotalAlloc = %.2f Mb", b2mb(m.TotalAlloc))
	s = s + fmt.Sprintf("\tHeapAlloc = %.2f Mb", b2mb(m.HeapAlloc))
	s = s + fmt.Sprintf("\tSys = %.2f Mb", b2mb(m.Sys))
	s = s + fmt.Sprintf("\tNumGC = %d\n", m.NumGC)
	return s
}

func main() {
	flag.Parse()
	s := printMemUsage()
	mem := []string{s}
	importRecords(*xmlfile, *jsonout)
	s = printMemUsage()
	mem = append(mem, s)
	fmt.Printf("%v\n", mem)
}
