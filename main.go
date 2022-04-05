package main

import (
	"github.com/joho/godotenv"
	"licenses/platform/middleware/fbauth"
	"licenses/platform/router"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	client, err := fbauth.InitAuth()

	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}

	rtr := router.New(client)

	log.Print("Server listening on http://localhost:8080/")
	if err := http.ListenAndServe("0.0.0.0:8080", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
