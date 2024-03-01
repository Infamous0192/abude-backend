package accounts

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/accounts/account"
	"abude-backend/internal/pkg/accounts/category"
)

func LoadRoutes(r *common.Router) {
	categoryService := category.NewService(r.DB)
	accountService := account.NewService(r.DB)

	categoryHandler := category.NewController(r.Controller, categoryService)
	r.Router.Get("/account-category", r.Auth(1), categoryHandler.All)
	r.Router.Get("/account-category/:id", r.Auth(1), categoryHandler.One)
	r.Router.Post("/account-category", r.Auth(2), categoryHandler.Create)
	r.Router.Put("/account-category/:id", r.Auth(2), categoryHandler.Update)
	r.Router.Delete("/account-category/:id", r.Auth(2), categoryHandler.Delete)

	accountHandler := account.NewController(r.Controller, accountService)
	r.Router.Get("/account", r.Auth(1), accountHandler.All)
	r.Router.Get("/account/:id", r.Auth(1), accountHandler.One)
	r.Router.Post("/account", r.Auth(2), accountHandler.Create)
	r.Router.Put("/account/:id", r.Auth(2), accountHandler.Update)
	r.Router.Delete("/account/:id", r.Auth(2), accountHandler.Delete)
}
