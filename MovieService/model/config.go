package model

type Config struct {
	MovieService Service `json:"MovieService"`
}

type Service struct {
	PortNumber int    `json:"port"`
	LogDir     string `json:"log_dir"`
	LogFile    string `json:"log_file"`
	MongoURL   string `json:"mongoUrl"`
}
