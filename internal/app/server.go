package app

import (
	"abude-backend/internal/config"
	"abude-backend/pkg/exception"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type ServerInstance struct {
	*fiber.App
	Config *config.ServerConfig
}

func (s *ServerInstance) Setup(config *config.ServerConfig) {
	s.Config = config
	s.LoadPath()

	s.App = fiber.New(fiber.Config{
		Concurrency:  256,
		ServerHeader: s.Config.Name,
		BodyLimit:    s.Config.UploadLimit,
		ErrorHandler: CustomErrorHandler,
	})

	s.LoadMiddlewares()
	s.LoadStatic()
	s.LoadSwagger()
}

func (s *ServerInstance) LoadMiddlewares() {
	s.App.Use(recover.New())
	s.App.Use(logger.New())
	s.App.Use(cors.New(config.CorsConfig))
}

func (s *ServerInstance) LoadStatic() {
	// Upload Path
	s.App.Static("/uploads", s.Config.UploadPath, fiber.Static{
		Compress:      true,
		ByteRange:     true,
		CacheDuration: 24 * time.Hour,
	})

	// Public Path
	s.App.Static("/", s.Config.PublicPath)

	// Redirect when access non-index route
	s.App.Use(func(c *fiber.Ctx) error {
		// Skip if paths starting with "/api" or "/docs"
		if strings.HasPrefix(c.Path(), "/api") || strings.HasPrefix(c.Path(), "/docs") {
			return c.Next()
		}

		// Serve index.html for all other non-file routes
		if _, err := os.Stat(filepath.Join(s.Config.PublicPath, c.Path())); os.IsNotExist(err) {
			return c.SendFile(filepath.Join(s.Config.PublicPath, "index.html"))
		}
		return c.Next()
	})
}

func (s *ServerInstance) LoadSwagger() {
	s.App.Get("/docs/*", swagger.New(swagger.Config{
		URL: fmt.Sprintf("%s/docs/doc.json", s.Config.Url),
	}))
}

func (s *ServerInstance) LoadPath() {
	if s.Config.Url == "" {
		s.Config.Url = fmt.Sprintf("http://%s:%s", s.Config.Host, s.Config.Port)
	}
	path, _ := os.Getwd()
	if s.Config.ExecPath {
		path = getPath()
	}
	s.Config.Path = path
	s.Config.UploadPath = MakeDir(filepath.Join(path, s.Config.UploadPath))
	s.Config.PublicPath = MakeDir(filepath.Join(path, s.Config.PublicPath))
	s.Config.UploadLimit = s.Config.UploadLimit * 1024 * 1024
}

func (s *ServerInstance) Serve(addr ...string) error {
	a := s.Config.Host + ":" + s.Config.Port
	if len(addr) != 0 {
		a = addr[0]
	}

	return s.Listen(a)
}

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	switch err := err.(type) {
	case exception.HttpError:
		return ctx.Status(err.Code).JSON(err)
	default:
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
}

func getPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func MakeDir(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	return path
}
