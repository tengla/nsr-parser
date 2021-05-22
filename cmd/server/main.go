package main

import (
	"log"
	"net/http"
)

func main() {
	docs, err := getStopPlaces()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", muxHandler(docs))
	log.Println("Starting server at :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
