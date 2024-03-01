package transactions

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/transactions/expense"
	"abude-backend/internal/pkg/transactions/purchase"
	"abude-backend/internal/pkg/transactions/sale"
	"abude-backend/internal/pkg/transactions/wage"
)

func LoadRoutes(r *common.Router) {
	saleService := sale.NewService(r.DB)
	purchaseService := purchase.NewService(r.DB)
	expenseService := expense.NewService(r.DB)
	wageService := wage.NewService(r.DB)

	saleHandler := sale.NewController(r.Controller, saleService)
	r.Router.Get("/sale", r.Auth(1), saleHandler.All)
	r.Router.Get("/sale/summary", r.Auth(1), saleHandler.GetSummary)
	r.Router.Get("/sale/:id", r.Auth(1), saleHandler.One)
	r.Router.Post("/sale", r.Auth(1), saleHandler.Create)
	r.Router.Put("/sale/:id", r.Auth(2), saleHandler.Update)
	r.Router.Delete("/sale/:id", r.Auth(2), saleHandler.Delete)
	r.Router.Patch("/sale/:id/cancel", r.Auth(1), saleHandler.Cancel)

	purchaseHandler := purchase.NewController(r.Controller, purchaseService)
	r.Router.Get("/purchase", r.Auth(1), purchaseHandler.All)
	r.Router.Get("/purchase/summary", r.Auth(1), purchaseHandler.GetSummary)
	r.Router.Get("/purchase/:id", r.Auth(1), purchaseHandler.One)
	r.Router.Post("/purchase", r.Auth(1), purchaseHandler.Create)
	r.Router.Put("/purchase/:id", r.Auth(2), purchaseHandler.Update)
	r.Router.Delete("/purchase/:id", r.Auth(2), purchaseHandler.Delete)
	r.Router.Patch("/purchase/:id/cancel", r.Auth(1), purchaseHandler.Cancel)

	expenseHandler := expense.NewController(r.Controller, expenseService)
	r.Router.Get("/expense", r.Auth(1), expenseHandler.All)
	r.Router.Get("/expense/summary", r.Auth(1), expenseHandler.GetSummary)
	r.Router.Get("/expense/:id", r.Auth(1), expenseHandler.One)
	r.Router.Post("/expense", r.Auth(1), expenseHandler.Create)
	r.Router.Put("/expense/:id", r.Auth(2), expenseHandler.Update)
	r.Router.Delete("/expense/:id", r.Auth(2), expenseHandler.Delete)
	r.Router.Patch("/expense/:id/cancel", r.Auth(1), expenseHandler.Cancel)

	wageHandler := wage.NewController(r.Controller, wageService)
	r.Router.Get("/wage", r.Auth(1), wageHandler.All)
	r.Router.Get("/wage/summary", r.Auth(1), wageHandler.GetSummary)
	r.Router.Get("/wage/:id", r.Auth(1), wageHandler.One)
	r.Router.Post("/wage", r.Auth(1), wageHandler.Create)
	r.Router.Put("/wage/:id", r.Auth(2), wageHandler.Update)
	r.Router.Delete("/wage/:id", r.Auth(2), wageHandler.Delete)
	r.Router.Patch("/wage/:id/cancel", r.Auth(1), wageHandler.Cancel)
}
