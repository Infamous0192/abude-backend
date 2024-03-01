package supplier

import "abude-backend/internal/common"

func LoadRoutes(r *common.Router) {
	supplierService := NewService(r.DB)
	supplierHandler := NewController(r.Controller, supplierService)

	r.Router.Get("/supplier", r.Auth(0), supplierHandler.All)
	r.Router.Get("/supplier/:id", r.Auth(0), supplierHandler.One)
	r.Router.Post("/supplier", r.Auth(1), supplierHandler.Create)
	r.Router.Put("/supplier/:id", r.Auth(1), supplierHandler.Update)
	r.Router.Delete("/supplier/:id", r.Auth(1), supplierHandler.Delete)
}
