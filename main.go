package main

import (
	"log"
	"lol-ranked-chatbot/routers"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "5000"
	}

	router := routers.InitializeRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
