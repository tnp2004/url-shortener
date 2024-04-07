package auth

import "time"

type (
	UserLoginReq struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}

	RefreshTokenReq struct {
		CredentialId string `json:"credential_id" form:"credential_id" validate:"required,max=64"`
		RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required,max=500"`
	}

	CredentialRes struct {
		Id           string    `json:"_id"`
		UserId       string    `json:"user_id"`
		RoleCode     int       `json:"role_code"`
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token" `
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	LogoutReq struct {
		CredentialId string `json:"credential_id" form:"credential_id" validate:"required,max=64"`
	}

	AuthResponseData struct {
		Credential *CredentialRes `json:"credential"`
	}
)