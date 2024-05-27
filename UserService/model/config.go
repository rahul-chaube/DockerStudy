package model

type Config struct {
	UserService Service `json:"UserService"`
}

type Service struct {
	PortNumber int    `json:"port"`
	LogDir     string `json:"log_dir"`
	LogFile    string `json:"log_file"`
}
