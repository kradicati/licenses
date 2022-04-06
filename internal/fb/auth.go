package fb

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/pkg/errors"
)

func InitAuth(app *firebase.App) (*auth.Client, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "error initializing fb auth (creating client)")
	}

	return client, nil
}
