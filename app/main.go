package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message  string `json:"message"`
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}` + "\n"))
}

func main() {
	version := os.Getenv("APP_VERSION")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		response := Response{
			Message:  "Hello",
			Version:  version,
			Hostname: hostname,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/health", health)

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		response := Response{
			Version:  version,
			Hostname: hostname,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Printf("Starting app, version %s", version)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}