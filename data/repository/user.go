package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"licenses/data/model"
)

type UserRepository struct {
	Client     *firestore.Client
	Collection *firestore.CollectionRef
}

func (repo UserRepository) Get(id string, obj *model.User) error {
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

func (repo UserRepository) Insert(id string, obj *model.User) error {
	_, err := repo.Collection.Doc(id).Create(context.Background(), *obj)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) Update(id string, update *[]firestore.Update) error {
	_, err := repo.Collection.Doc(id).Update(context.Background(), *update)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) UpdateApiKey(id string, apiKey string) error {
	update := []firestore.Update{{Path: "api_key", Value: apiKey}}
	err := repo.Update(id, &update)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) Delete(id string) error {
	_, err := repo.Collection.Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) FindAll(query *firestore.Query, obj *[]model.User) error {
	iter := query.Documents(context.Background())

	err := iteratorToArray(iter, obj)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) FindAllByApiKey(apiKey string, obj *[]model.User) error {
	query := repo.Collection.Where("api_key", "==", apiKey)
	err := repo.FindAll(&query, obj)

	return err
}

func NewUserRepository(client *firestore.Client) *UserRepository {
	repo := &UserRepository{
		Client:     client,
		Collection: client.Collection("users"),
	}

	return repo
}
