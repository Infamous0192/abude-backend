package app

import (
	"abude-backend/internal/common"
	"abude-backend/internal/config"
	"abude-backend/pkg/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Creds common.Creds
}

type JwtInstance struct {
	Config *config.JwtConfig
}

func (cfg *JwtInstance) Setup(config *config.JwtConfig) {
	cfg.Config = config
}

func (cfg *JwtInstance) Hash(creds common.Creds) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Creds: creds,
	})

	tokenString, err := token.SignedString([]byte(cfg.Config.Secret))
	if err != nil {
		return "", exception.Http(fiber.StatusInternalServerError, err.Error())
	}

	return tokenString, nil
}

func (cfg *JwtInstance) Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Config.Secret), nil
	})

	if err != nil {
		return nil, exception.Http(fiber.StatusUnauthorized, "Token invalid or expired")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, exception.Http(fiber.StatusUnauthorized, "Token invalid or expired")
	}

	return claims, nil
}
