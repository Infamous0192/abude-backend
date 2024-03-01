package company

import "abude-backend/internal/common"

func LoadRoutes(r *common.Router) {
	companyService := NewService(r.DB)
	companyHandler := NewController(r.Controller, companyService)

	r.Router.Get("/company", r.Auth(1), companyHandler.All)
	r.Router.Get("/company/:id", r.Auth(2), companyHandler.One)
	r.Router.Post("/company", r.Auth(2), companyHandler.Create)
	r.Router.Put("/company/:id", r.Auth(2), companyHandler.Update)
	r.Router.Delete("/company/:id", r.Auth(2), companyHandler.Delete)
	r.Router.Get("/company/:id/owner", r.Auth(1), companyHandler.GetOwners)
	r.Router.Put("/company/:companyId/owner/:userId", r.Auth(1), companyHandler.AddOwner)
	r.Router.Delete("/company/:companyId/owner/:userId", r.Auth(1), companyHandler.RemoveOwner)
}
