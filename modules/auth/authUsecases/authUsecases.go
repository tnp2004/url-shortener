package authUsecases

import (
	"github.com/tnp2004/url-shortener/modules/auth/authRepositories"
)

type (
	IAuthUsecases interface{}

	authUsecases struct {
		authRepository authRepositories.IAuthRepository
	}
)

func NewAuthUsecases(authRepository authRepositories.IAuthRepository) IAuthUsecases {
	return &authUsecases{authRepository: authRepository}
}
