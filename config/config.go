package config

import (
	"os"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Stage   string
		Name    string
		Port    string
		Version string
	}

	Db struct {
		Url string
	}
)

func LoadConfig() Config {
	return Config{
		App: App{
			Stage:   os.Getenv("APP_STAGE"),
			Name:    os.Getenv("APP_NAME"),
			Port:    "0.0.0.0:" + os.Getenv("PORT"),
			Version: os.Getenv("VERSION"),
		},
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
	}
}
