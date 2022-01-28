package main

import (
	"log"
	"net/http"
)

func main() {
	places, err := GetStopPlaces()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", MuxHandler(places))
	log.Println("Starting server at :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
