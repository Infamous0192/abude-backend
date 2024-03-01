package employee

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/user"
)

func LoadRoutes(r *common.Router) {
	userService := user.NewService(r.DB)
	employeeService := NewService(r.DB, userService)
	employeeHandler := NewController(r.Controller, employeeService)

	r.Router.Get("/employee", r.Auth(1), employeeHandler.All)
	r.Router.Get("/employee/:id", r.Auth(0), employeeHandler.One)
	r.Router.Post("/employee", r.Auth(1), employeeHandler.Create)
	r.Router.Put("/employee/:id", r.Auth(1), employeeHandler.Update)
	r.Router.Delete("/employee/:id", r.Auth(1), employeeHandler.Delete)
}
