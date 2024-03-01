package user

import "abude-backend/internal/common"

func LoadRoutes(r *common.Router) {
	userService := NewService(r.DB)
	userHandler := NewController(r.Controller, userService)

	r.Router.Get("/user", r.Auth(1), userHandler.All)
	r.Router.Get("/user/:id", r.Auth(2), userHandler.One)
	r.Router.Post("/user", r.Auth(2), userHandler.Create)
	r.Router.Put("/user/:id", r.Auth(2), userHandler.Update)
	r.Router.Delete("/user/:id", r.Auth(2), userHandler.Delete)
}
