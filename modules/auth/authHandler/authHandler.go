package authHandler

import "github.com/tnp2004/url-shortener/modules/auth/authUsecases"

type (
	IAuthHandler interface{}

	authHandler struct {
		authUsecases authUsecases.IAuthUsecases
	}
)

func NewAuthHandler(authUsecases authUsecases.IAuthUsecases) IAuthHandler {
	return &authHandler{authUsecases: authUsecases}
}
