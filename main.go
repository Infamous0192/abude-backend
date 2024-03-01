package main

import (
	_ "abude-backend/docs"
	"abude-backend/internal/app"
	"abude-backend/internal/database/migrations"
	"abude-backend/internal/database/seeders"
	"flag"
	"log"
)

// @title Abude Backend
// @version 2.0.0
// @description Aplikasi Backend (API) untuk project Abude
// @termsOfService http://swagger.io/terms/
// @contact.name Dwa Meizadewa
// @contact.email infamous0192@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @schemes https http
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description Authorization for JWT
func main() {
	configFile := flag.String("config", "config.yml", "User Config file from user")
	migrate := flag.Bool("migrate", false, "Auto migrate")
	seed := flag.Bool("seed", false, "Seeder")

	flag.Parse()
	app := app.Load(*configFile)

	if *seed {
		seeders.DatabaseSeeder(app.Database.DB)
		return
	}

	if *migrate {
		migrations.AutoMigrate(app.Database.DB)
	}

	log.Fatal(app.Server.Serve())
}
