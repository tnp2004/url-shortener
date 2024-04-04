package converterHandler

import (
	"context"
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
		SearchDestination(c *fiber.Ctx) error
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
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(converter.ConverterReq)
	if err := wrapper.Bind(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.converterUsecase.GetShortUrl(ctx, req)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, res)
}

func (h *converterHandler) SearchDestination(c *fiber.Ctx) error {
	ctx := context.Background()

	shortId := c.Params("short_id")

	des, err := h.converterUsecase.SearchDestination(ctx, &converter.SearchShortIdReq{
		ShortId: shortId,
	})
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	return response.Redirect(c, des)
}
