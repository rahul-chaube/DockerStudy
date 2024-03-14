package utility

type Config struct {
	UserService Service `json:"UserService"`
}

type Service struct {
	PortNumber string `json:"port"`
}
