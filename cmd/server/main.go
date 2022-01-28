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
	log.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
