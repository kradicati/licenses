package model

type User struct {
	Id       string   `firestore:"id"`
	Licenses []string `firestore:"licenses"`
	Roles    []string `firestore:"roles"`
	ApiKey   string   `firestore:"api_key"`
}
