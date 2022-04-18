package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"licenses/internal/fb"
	"licenses/platform/router"
	"log"
	"net/http"
	"os"
)

//TODO put everything as {"message":..."} for errors
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load the env vars: %v", err)
	}

	if os.Getenv("KEYS") == "true" {
		err := handleKeys()

		if err != nil {
			log.Fatalf("failed to load keys: %v", err)
			return
		}
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

	store, err := fb.InitFireStore(app)
	if err != nil {
		log.Fatalln("failed to init firestore", err)
	}

	rtr := router.New(auth, store)

	port := os.Getenv("PORT")

	log.Print("Server listening on http://localhost:" + port + "/")
	if err := http.ListenAndServe("0.0.0.0:"+port, rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}

func handleKeys() error {
	private, err := getPrivateKey()
	if err != nil {
		return err
	}

	router.PrivateKey = private

	pub, err := base64.StdEncoding.DecodeString(os.Getenv("PUBLIC_KEY"))

	if err != nil {
		return err
	}

	router.PublicKey = pub

	return nil
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	key, err := base64.StdEncoding.DecodeString(os.Getenv("PRIVATE_KEY"))

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(key)
	der, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return der, err
}
