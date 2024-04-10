package main

import (
	"context"

	"github.com/tnp2004/url-shortener/config"
	"github.com/tnp2004/url-shortener/pkg/database"
	"github.com/tnp2004/url-shortener/server"
)

func main() {
	cfg := config.LoadConfig()

	ctx := context.Background()

	db := database.DbConnect(ctx, &cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, &cfg, db)
}
