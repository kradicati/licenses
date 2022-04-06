package main

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"licenses/internal/fb"
	"licenses/platform/router"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CONFIG_FILE"))

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln("failed to firebase app", err)
	}

	auth, err := fb.InitAuth(app)
	if err != nil {
		log.Fatalln("failed to init firebase auth", err)
	}

	_, err = fb.InitFireStore(app)
	if err != nil {
		log.Fatalln("failed to init firestore", err)
	}

	rtr := router.New(auth)

	port := os.Getenv("PORT")

	log.Print("Server listening on http://localhost:" + port + "/")
	if err := http.ListenAndServe("0.0.0.0:"+port, rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
