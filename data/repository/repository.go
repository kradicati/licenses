package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"licenses/data/model"
)

type Repository[T interface{}] interface {
	Get(id string, obj *T) T
	Insert(id string, obj *T)
}

type LicenseRepository struct {
	Client     *firestore.Client
	Collection *firestore.CollectionRef
}

func (repo LicenseRepository) Get(id string, obj *model.License) error {
	get, err := repo.Collection.Doc(id).Get(context.Background())
	if err != nil {
		return err
	}

	err = get.DataTo(obj)
	if err != nil {
		return err
	}
	return nil
}

func (repo LicenseRepository) Insert(id string, obj *model.License) error {
	return nil
}
