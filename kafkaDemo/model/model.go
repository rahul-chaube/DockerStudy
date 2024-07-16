package model

type Order struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	UserId  string `json:"userId"`
	Ammount int64  `json:"ammount"`
}
