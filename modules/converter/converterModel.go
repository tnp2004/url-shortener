package converter

type (
	ConverterReq struct {
		URL string `json:"url" validate:"required,url"`
	}

	ConverterRes struct {
		ShortenedURL string `json:"shortened_url"`
	}

	SearchShortIdReq struct {
		ShortId string
	}
)
