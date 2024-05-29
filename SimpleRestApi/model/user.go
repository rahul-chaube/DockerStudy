package model

var UserData []User

type User struct {
	Id   string
	Name string `json:"name"`
}
