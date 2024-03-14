package model

import "time"

type User struct {
	Id         string
	Name       string
	Age        string
	AddedAt    time.Time
	LastUpdate time.Time
}
