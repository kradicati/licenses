package model

type User struct {
	Id       string
	UserName string
	Licenses []string
	Roles    []string
}
