package model

type Favourate struct {
	Id        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Url       string `json:"url" bson:"url"`
	HomeWorld string `json:"homeworld" bson:"homeworld"`
}
