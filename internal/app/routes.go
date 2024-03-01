package app

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/accounts"
	"abude-backend/internal/pkg/attendances"
	"abude-backend/internal/pkg/auth"
	"abude-backend/internal/pkg/company"
	"abude-backend/internal/pkg/employee"
	"abude-backend/internal/pkg/handover"
	"abude-backend/internal/pkg/inventories"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/internal/pkg/supplier"
	"abude-backend/internal/pkg/transactions"
	"abude-backend/internal/pkg/turnover"
	"abude-backend/internal/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *AppInstance) {
	web := app.Server.App
	db := app.Database.DB

	api := web.Group("/api")
	authService := auth.NewService(db, app.Jwt.Config)
	authMiddleware := auth.NewMiddleware(authService)

	ctrl := &common.BaseController{
		Validation: &app.Validation,
		DB:         db,
	}

	router := &common.Router{
		Router:     api,
		Controller: ctrl,
		DB:         db,
		Auth:       authMiddleware,
	}

	user.LoadRoutes(router)
	auth.LoadRoutes(router, authService)
	employee.LoadRoutes(router)
	company.LoadRoutes(router)
	outlet.LoadRoutes(router)
	supplier.LoadRoutes(router)
	inventories.LoadRoutes(router)
	accounts.LoadRoutes(router)
	transactions.LoadRoutes(router)
	handover.LoadRoutes(router)
	turnover.LoadRoutes(router)
	attendances.LoadRoutes(router)

	web.All("/api/*", func(ctx *fiber.Ctx) error {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Route not found",
		})
	})
}
