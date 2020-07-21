package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Mysql struct {
	Ip       string `json:"Ip" `
	Port     string `json:"Port" `
	User     string `json:"User" `
	Password string `json:"Password" `
	Db       string `json:"Db" `
}

type Redis struct {
	Ip   string `json:"Ip"`
	Port string `json:"Port"`
	Auth string `json:"Auth"`
}

type MyConfig struct {
	Mysql
	Redis
}

var Config = MyConfig{
	Mysql: Mysql{
		Ip:       os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Db:       os.Getenv("DB_NAME"),
	},
	Redis: Redis{
		Ip:   os.Getenv("REDIS_ENDPOINT"),
		Port: os.Getenv("REDIS_PORT"),
		Auth: os.Getenv("REDIS_AUTH"),
	},
}
