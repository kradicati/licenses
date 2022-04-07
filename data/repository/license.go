package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"licenses/data/model"
)

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
	_, err := repo.Collection.Doc(id).Create(context.Background(), *obj)
	if err != nil {
		return err
	}

	return nil
}

func (repo LicenseRepository) FindAll(query *firestore.Query, obj *[]model.License) error {
	iter := query.Documents(context.Background())

	err := iteratorToArray(iter, obj)
	if err != nil {
		return err
	}

	return nil
}

func (repo LicenseRepository) FindAllByIp(ip string, obj *[]model.License) error {
	query := repo.Collection.Where("ip_log", "array-contains", ip)
	err := repo.FindAll(&query, obj)

	return err
}

func NewLicenseRepository(client *firestore.Client) *LicenseRepository {
	repo := &LicenseRepository{
		Client:     client,
		Collection: client.Collection("licenses"),
	}

	return repo
}