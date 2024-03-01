package app

import (
	"abude-backend/internal/config"
	"abude-backend/pkg/validation"
)

type AppInstance struct {
	Server   ServerInstance
	Database DatabaseInstance
	Jwt      JwtInstance

	// Plugins
	Validation validation.Validation
}

func Load(configFile string) *AppInstance {
	config := config.Load(configFile)

	// Database.Setup(&app.Database)
	// Server.Setup(&app.Server)
	// Jwt.Setup(&app.Jwt)

	// // Load plugins
	// Validation.Setup(Database.DB)

	app := new(AppInstance)
	app.Database.Setup(&config.Database)
	app.Server.Setup(&config.Server)
	app.Jwt.Setup(&config.Jwt)

	app.Validation.Setup(app.Database.DB)

	LoadRoutes(app)

	return app
}
