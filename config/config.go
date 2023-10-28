package config

type MySQL struct {
	Host     string
	DbName   string
	Port     int
	username string
	Password string
	Charset  string
}

type Redis struct {
	Host string
	Port string
}
