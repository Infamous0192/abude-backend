package attendances

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/attendances/shift"
)

func LoadRoutes(r *common.Router) {
	shiftService := shift.NewService(r.DB)
	shiftHandler := shift.NewController(r.Controller, shiftService)

	r.Router.Get("/shift", r.Auth(1), shiftHandler.All)
	r.Router.Get("/shift/:id", r.Auth(1), shiftHandler.One)
	r.Router.Post("/shift", r.Auth(2), shiftHandler.Create)
	r.Router.Put("/shift/:id", r.Auth(2), shiftHandler.Update)
	r.Router.Delete("/shift/:id", r.Auth(2), shiftHandler.Delete)
}
