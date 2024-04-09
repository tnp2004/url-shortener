package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Stage   string
		Name    string
		Url     string
		Version string
	}

	Db struct {
		Url string
	}
)

func LoadConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		App: App{
			Stage:   os.Getenv("APP_STAGE"),
			Name:    os.Getenv("APP_NAME"),
			Url:     os.Getenv("APP_URL"),
			Version: os.Getenv("VERSION"),
		},
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
	}
}
