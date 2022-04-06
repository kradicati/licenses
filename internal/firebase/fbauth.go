package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"os"
)

func InitAuth() (*auth.Client, error) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CONFIG_FILE"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing firebase auth (create firebase app)")
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "error initializing firebase auth (creating client)")
	}

	return client, nil
}
