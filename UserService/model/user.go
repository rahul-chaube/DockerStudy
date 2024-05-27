package model

import "time"

var UserData []User

type User struct {
	Id         string
	Name       string
	Age        string
	AddedAt    time.Time
	LastUpdate time.Time
}
