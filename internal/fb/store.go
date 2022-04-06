package fb

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
)

func InitFireStore(app *firebase.App) (*firestore.Client, error) {
	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "error initializing fb auth (creating client)")
	}

	return client, nil
}
