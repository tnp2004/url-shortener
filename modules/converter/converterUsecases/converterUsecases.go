package converterUsecases

import "github.com/tnp2004/url-shortener/modules/converter/converterRepositories"

type (
	IConverterUsecases interface{}
	converterUsecases  struct {
		converterRepository converterRepositories.IConverterRepository
	}
)
