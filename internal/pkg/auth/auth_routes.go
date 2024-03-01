package auth

import (
	"abude-backend/internal/common"
)

func LoadRoutes(r *common.Router, authService *AuthService) {
	authHandler := NewController(r.Controller, authService)

	r.Router.Post("/auth/login", authHandler.Login)
	r.Router.All("/auth/me", r.Auth(0), authHandler.Verify)
}
