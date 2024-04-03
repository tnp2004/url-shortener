package server

import "github.com/gofiber/fiber/v2"

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
	m.router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
