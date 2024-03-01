package common

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Router struct {
	Router     fiber.Router
	Controller *BaseController
	DB         *gorm.DB
	Auth       func(level int) func(*fiber.Ctx) error
}
