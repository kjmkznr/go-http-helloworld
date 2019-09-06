package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	msg := os.Getenv("MESSAGE")
	if msg == "" {
		msg = "Hello, World from " + getHostname()
	}
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(msg))
	})))
}

func getHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
