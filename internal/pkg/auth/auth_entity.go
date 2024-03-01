package auth

import (
	"abude-backend/internal/common"

	"github.com/golang-jwt/jwt/v5"
)

const (
	CredsKey = "creds"
)

type LoginDTO struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type Authenticated struct {
	Creds common.Creds `json:"creds"`
	Token string       `json:"token"`
}

type Claims struct {
	jwt.RegisteredClaims
	common.Creds
}
