package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/tengla/nsr-parser/stopplace"
)

type payload map[string]string

func (p payload) hasEvery(keys ...string) bool {
	for _, k := range keys {
		if !p.has(k) {
			return false
		}
	}
	return true
}

func (p payload) has(key string) bool {
	return len(p[key]) > 0
}

func getStopPlaces() ([]stopplace.StopPlace, error) {
	file, err := os.Open("./nsr.current.json")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	docs := make([]stopplace.StopPlace, 0)
	n := 0
	for scanner.Scan() {
		bytes := scanner.Bytes()
		var v stopplace.StopPlace
		json.Unmarshal(bytes, &v)
		docs = append(docs, v)
		n += 1
		fmt.Printf("\rParsing progress %d", n)
	}
	fmt.Printf("\rDone parsing %d docs\n", len(docs))
	return docs, nil
}

func isPath(r *http.Request, method string, path string) bool {
	return r.Method == method && r.URL.Path == path
}

func isPathMatch(r *http.Request, method string, path string) bool {
	fmt.Printf("%s %s\n", r.URL.Path, path)
	re, _ := regexp.Compile(path)
	return r.Method == method && re.MatchString(r.URL.Path)
}

func main() {
	docs, err := getStopPlaces()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		rw.Header().Add("Content-Type", "application/json; charset=utf-8")
		var p payload
		json.NewDecoder(r.Body).Decode(&p)
		switch true {
		case isPathMatch(r, "GET", "/NSR:StopPlace:[0-9]+"):
			segments := strings.Split(r.URL.Path, "/")
			id := segments[len(segments)-1]
			var stop stopplace.StopPlace
			for _, s := range docs {
				if s.ID == id {
					stop = s
				}
			}
			json.NewEncoder(rw).Encode(stop)
		case isPath(r, "POST", "/name") && p.has("Value"):
			stops := make([]stopplace.StopPlace, 0)
			re, _ := regexp.Compile("(?i)" + p["Value"])
			for _, stop := range docs {
				if re.MatchString(stop.Name.Text) {
					stops = append(stops, stop)
				}
			}
			json.NewEncoder(rw).Encode(stops)
		case isPath(r, "POST", "/keylist") && p.hasEvery("Key", "Value"):
			stops := make([]stopplace.StopPlace, 0)
			re, err := regexp.Compile(p["Value"])
			if err != nil {
				http.Error(rw, "Not allowed", http.StatusMethodNotAllowed)
			} else {
				for _, stop := range docs {
					for _, kv := range stop.KeyList.KeyValue {
						if kv.Key == p["Key"] && re.MatchString(kv.Value) {
							stops = append(stops, stop)
						}
					}
				}
			}
			json.NewEncoder(rw).Encode(stops)
		default:
			message := map[string]string{
				"/keylist":              "method: POST, payload: {\"Key\":\"key\",\"Value\":\"value\"}",
				"/name":                 "method: POST, payload: {\"Value\":\"Oslo S\"}",
				"/NSR:StopPlace:[0-9]+": "method: GET",
			}
			json.NewEncoder(rw).Encode(message)
		}
	})
	log.Println("Starting server at :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
