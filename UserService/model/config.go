package model

type Config struct {
	UserService Service `json:"UserService"`
}

type Service struct {
	PortNumber int `json:"port"`
}
