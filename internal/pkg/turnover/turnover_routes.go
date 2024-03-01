package turnover

import "abude-backend/internal/common"

func LoadRoutes(r *common.Router) {
	turnoverService := NewService(r.DB)
	turnoverHandler := NewController(r.Controller, turnoverService)

	r.Router.Get("/turnover", r.Auth(0), turnoverHandler.All)
	r.Router.Get("/turnover/:id", r.Auth(0), turnoverHandler.One)
	r.Router.Post("/turnover", r.Auth(1), turnoverHandler.Create)
	r.Router.Put("/turnover/:id", r.Auth(1), turnoverHandler.Update)
	r.Router.Delete("/turnover/:id", r.Auth(1), turnoverHandler.Delete)
}
