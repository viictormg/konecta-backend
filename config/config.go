package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host       string
	UserDB     string
	PasswordDB string
	HostDB     string
	NameDB     string
	JwtSecret  string
}

func NewConfig() (Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return Config{}, err
	}

	return Config{
		Host:       os.Getenv("HTTP_ADDR"),
		UserDB:     os.Getenv("USER_DB"),
		PasswordDB: os.Getenv("PASSWORD_DB"),
		HostDB:     os.Getenv("HOST_DB"),
		NameDB:     os.Getenv("NAME_DB"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
	}, nil
}
