package model

type User struct {
	Id       string   `firestore:"id"`
	UserName string   `firestore:"user_name"`
	Licenses []string `firestore:"licenses"`
	Roles    []string `firestore:"roles"`
}
