package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bar/id", func(rw http.ResponseWriter, r *http.Request) {
		message := map[string]string{
			"message": "Hello",
		}
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(message)
	})
	log.Println("Starting server at :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
