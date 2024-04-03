package converterHandler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tnp2004/url-shortener/modules/converter"
	"github.com/tnp2004/url-shortener/modules/converter/converterUsecases"
	"github.com/tnp2004/url-shortener/pkg/request"
	"github.com/tnp2004/url-shortener/pkg/response"
)

type (
	IConverterHandler interface {
		Greeting(c *fiber.Ctx) error
		Convert(c *fiber.Ctx) error
	}
	converterHandler struct {
		converterUsecase converterUsecases.IConverterUsecases
	}
)

func NewConverterHandler(usecase converterUsecases.IConverterUsecases) IConverterHandler {
	return &converterHandler{converterUsecase: usecase}
}

func (h *converterHandler) Greeting(c *fiber.Ctx) error {
	return c.SendString("Hi guys this is url shortener!")
}

func (h *converterHandler) Convert(c *fiber.Ctx) error {
	wrapper := request.ContextWrapper(c)

	req := new(converter.ConverterReq)

	if err := wrapper.Bind(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	return response.Success(c, http.StatusOK, req)
}
