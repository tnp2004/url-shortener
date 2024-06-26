package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnp2004/url-shortener/modules/converter/converterHandler"
	"github.com/tnp2004/url-shortener/modules/converter/converterRepositories"
	"github.com/tnp2004/url-shortener/modules/converter/converterUsecases"
)

type (
	IModuleFactory interface {
		ConverterModule()
	}

	moduleFactory struct {
		router fiber.Router
		server *server
	}
)

func InitModules(r fiber.Router, s *server) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
	}
}

func (m *moduleFactory) ConverterModule() {
	repository := converterRepositories.NewConverterRepository(m.server.db)
	usecase := converterUsecases.NewConverterUsecases(repository, &m.server.cfg.App)
	handler := converterHandler.NewConverterHandler(usecase)

	router := m.router

	router.Get("/healthcheck", handler.HealthCheck)
	router.Get("/:short_id", handler.SearchDestination)

	router.Post("/convert", handler.Convert)
}
