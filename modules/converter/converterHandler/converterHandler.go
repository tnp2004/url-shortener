package converterHandler

import "github.com/tnp2004/url-shortener/modules/converter/converterUsecases"

type (
	IConverterHandler interface{}
	converterHandler  struct {
		converterUsecase converterUsecases.IConverterUsecases
	}
)
