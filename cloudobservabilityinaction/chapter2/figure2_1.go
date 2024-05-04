package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func handleEcho(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	log.WithFields(log.Fields{
		"service": "echo",
		}).Info("Got input: ", message)
	if rand.Intn(100) > 90 {
		log.WithFields(log.Fields{
			"service": "echo",
			}).Error("Something really bad happened")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, message)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFormatter(&log.JSONFormatter{})
	http.HandleFunc("/echo", handleEcho)
	log.Fatal(http.ListenAndServe(":4242", nil))
}
