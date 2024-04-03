package main

import (
	"context"
	"log"
	"os"

	"github.com/tnp2004/url-shortener/config"
	"github.com/tnp2004/url-shortener/pkg/database"
	"github.com/tnp2004/url-shortener/server"
)

func main() {
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		} else if len(os.Args) > 2 {
			log.Fatal("Error: too many arguments")
		}

		return os.Args[1]
	}())

	ctx := context.Background()

	db := database.DbConnect(ctx, &cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, &cfg, db)
}
