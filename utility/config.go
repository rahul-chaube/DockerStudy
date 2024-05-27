package utility

type Config struct {
	UserService Service `json:"UserService"`
}

type Service struct {
	PortNumber string `json:"port"`
	LogDir     string `json:"log_dir"`
	LogFile    string `json:"log_file"`
}
