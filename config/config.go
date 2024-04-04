package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App  App
		Db   Db
		Jwt  Jwt
		Grpc Grpc
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

	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSecretKey     string
		AccessDuration   int64
		RefreshDuration  int64
	}

	Grpc struct {
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
		Jwt: Jwt{
			AccessSecretKey:  os.Getenv("JWT_ACCESS_SECRET_KEY"),
			RefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),
			ApiSecretKey:     os.Getenv("JWT_API_SECRET_KEY"),
			AccessDuration: func() int64 {
				duration, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error: load access duration failed")
				}
				return duration
			}(),
			RefreshDuration: func() int64 {
				duration, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_DURATION"), 10, 64)
				if err != nil {
					log.Fatal("Error: load refresh duration failed")
				}
				return duration
			}(),
		},
		Grpc: Grpc{
			Url: os.Getenv("GRPC_URL"),
		},
	}
}
