package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tengla/nsr-parser/stopplace"
)

// GetStopPlaces -talks for itself, no?
func GetStopPlaces() ([]stopplace.StopPlace, error) {
	file, err := os.Open("./nsr.current.json")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	places := make([]stopplace.StopPlace, 0)
	n := 0
	for scanner.Scan() {
		bytes := scanner.Bytes()
		var v stopplace.StopPlace
		json.Unmarshal(bytes, &v)
		places = append(places, v)
		n += 1
		fmt.Printf("\rParsing progress %d", n)
	}
	fmt.Printf("\rDone parsing %d docs\n", len(places))
	return places, nil
}

func isPath(r *http.Request, method string, path string) bool {
	return r.Method == method && r.URL.Path == path
}

// MuxHandler
func MuxHandler(places []stopplace.StopPlace) func(rw http.ResponseWriter, r *http.Request) {

	return func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		rw.Header().Add("Content-Type", "application/json; charset=utf-8")
		var search_params stopplace.SearchParams
		json.NewDecoder(r.Body).Decode(&search_params)
		switch true {
		case isPath(r, "POST", "/"):
			params := search_params.ExtractParams()
			fmt.Printf("%+v\n", params)
			encoder := json.NewEncoder(rw)
			for _, place := range places {
				if stopplace.Match(params, place) {
					encoder.Encode(place)
				}
			}
		default:
			message := map[string]string{
				"message": "Hello World",
			}
			json.NewEncoder(rw).Encode(message)
		}
	}
}
