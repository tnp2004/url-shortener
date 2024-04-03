package request

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	IContextWrapper interface {
		Bind(data any) error
	}

	contextWrapper struct {
		Context   *fiber.Ctx
		validator *validator.Validate
	}
)

func ContextWrapper(ctx *fiber.Ctx) IContextWrapper {
	return &contextWrapper{
		Context:   ctx,
		validator: validator.New(),
	}
}

func (c *contextWrapper) Bind(data any) error {
	if err := c.Context.BodyParser(data); err != nil {
		log.Printf("Error: Bind data failed: %s", err.Error())
		return errors.New("error: bind data failed")
	}

	if err := c.validator.Struct(data); err != nil {
		log.Printf("Error: Validate data failed: %s", err.Error())
		return errors.New("error: validate data failed")
	}

	return nil
}
