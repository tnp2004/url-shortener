package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tnp2004/url-shortener/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app *fiber.App
		cfg *config.Config
		db  *mongo.Client
	}
)

func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{
		app: fiber.New(fiber.Config{
			CaseSensitive: true,
		}),
		cfg: cfg,
		db:  db,
	}

	router := s.app.Group("/v1")

	modules := InitModules(router, s)
	modules.ConverterModule()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefulShutdown(pctx, quit)

	s.httpListening()
}

func (s *server) httpListening() {
	if err := s.app.Listen(s.cfg.App.Port); err != nil {
		log.Fatal("Error: ", err.Error())
	}
}

func (s *server) gracefulShutdown(pctx context.Context, quit <-chan os.Signal) {
	<-quit

	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := s.app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
