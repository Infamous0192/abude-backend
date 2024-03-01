package handover

import (
	"abude-backend/internal/common"
)

func LoadRoutes(r *common.Router) {
	handoverService := NewHandoverService(r.DB)
	proofService := NewProofService(r.DB)

	proofHandler := NewProofController(r.Controller, proofService)
	r.Router.Get("/handover/proof", r.Auth(1), proofHandler.All)
	r.Router.Get("/handover/proof/:id", r.Auth(0), proofHandler.One)
	r.Router.Post("/handover/proof", r.Auth(1), proofHandler.Create)
	r.Router.Put("/handover/proof/:id", r.Auth(1), proofHandler.Update)
	r.Router.Delete("/handover/proof/:id", r.Auth(1), proofHandler.Delete)

	handoverHandler := NewHandoverController(r.Controller, handoverService)
	r.Router.Get("/handover", r.Auth(1), handoverHandler.All)
	r.Router.Get("/handover/:id", r.Auth(0), handoverHandler.One)
	r.Router.Post("/handover", r.Auth(1), handoverHandler.Create)
	r.Router.Put("/handover/:id", r.Auth(1), handoverHandler.Update)
	r.Router.Delete("/handover/:id", r.Auth(1), handoverHandler.Delete)
}
