package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now()
	message := fmt.Sprintf("Pong!\n Time:%s. \nThis HTTP triggered function executed successfully.\n ", tm)
	fmt.Fprint(w, message)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("BROKER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/ping", pingHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
