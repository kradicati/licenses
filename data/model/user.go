package model

type User struct {
	Id       string   `firestore:"id"`
	Email    string   `firestore:"email"`
	Licenses []string `firestore:"licenses"`
	Roles    []string `firestore:"roles"`
	ApiKey   string   `firestore:"api_key"`
}
