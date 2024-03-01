package inventories

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/inventories/category"
	"abude-backend/internal/pkg/inventories/inventory"
	"abude-backend/internal/pkg/inventories/product"
)

func LoadRoutes(r *common.Router) {
	categoryService := category.NewService(r.DB)
	productService := product.NewService(r.DB)
	inventoryService := inventory.NewService(r.DB)

	categoryHandler := category.NewController(r.Controller, categoryService)
	r.Router.Get("/category", r.Auth(1), categoryHandler.All)
	r.Router.Get("/category/:id", r.Auth(1), categoryHandler.One)
	r.Router.Post("/category", r.Auth(1), categoryHandler.Create)
	r.Router.Put("/category/:id", r.Auth(1), categoryHandler.Update)
	r.Router.Delete("/category/:id", r.Auth(1), categoryHandler.Delete)

	productHandler := product.NewController(r.Controller, productService)
	r.Router.Get("/product", r.Auth(1), productHandler.All)
	r.Router.Get("/product/:id", r.Auth(1), productHandler.One)
	r.Router.Post("/product", r.Auth(1), productHandler.Create)
	r.Router.Put("/product/:id", r.Auth(1), productHandler.Update)
	r.Router.Delete("/product/:id", r.Auth(1), productHandler.Delete)

	inventoryHandler := inventory.NewController(r.Controller, inventoryService)
	r.Router.Get("/inventory/stock", r.Auth(1), inventoryHandler.GetStock)
	r.Router.Get("/inventory/summary", r.Auth(1), inventoryHandler.GetStockSummary)

	r.Router.Get("/inventory/recapitulation", r.Auth(1), inventoryHandler.GetRecaps)
	r.Router.Get("/inventory/recapitulation/:id", r.Auth(1), inventoryHandler.GetRecap)
	r.Router.Post("/inventory/recapitulation", r.Auth(1), inventoryHandler.CreateRecap)

	r.Router.Get("/inventory", r.Auth(1), inventoryHandler.All)
	r.Router.Get("/inventory/:id", r.Auth(1), inventoryHandler.One)
	r.Router.Put("/inventory", r.Auth(1), inventoryHandler.Add)
	r.Router.Delete("/inventory/:id", r.Auth(1), inventoryHandler.Delete)
}
