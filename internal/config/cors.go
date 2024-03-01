package config

import "github.com/gofiber/fiber/v2/middleware/cors"

var CorsConfig = cors.Config{
	AllowOrigins:     "*",
	AllowHeaders:     "Origin, Content-Type, AcceptX-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token, X-Tenant-Token",
	AllowMethods:     "POST, GET, OPTIONS, PATCH, PUT, DELETE, UPDATE",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
	MaxAge:           86400,
}
