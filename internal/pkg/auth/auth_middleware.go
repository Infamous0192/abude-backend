package auth

import (
	"abude-backend/internal/pkg/user"
	"abude-backend/pkg/exception"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware func(level int) func(*fiber.Ctx) error

func NewMiddleware(s *AuthService) AuthMiddleware {
	return func(level int) fiber.Handler {
		return func(ctx *fiber.Ctx) error {
			parts := strings.Split(ctx.Get(fiber.HeaderAuthorization), " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return exception.Http(fiber.StatusUnauthorized, "Unauthorized")
			}
			token := parts[1]

			claims, err := s.Parse(token)
			if err != nil {
				return err
			}

			if !isPermitted(claims.Role, level) {
				return exception.Forbidden()
			}

			ctx.Locals(CredsKey, claims.Creds)

			return ctx.Next()
		}
	}
}

func isPermitted(role string, level int) bool {
	switch role {
	case user.RoleSuperadmin:
		return true
	case user.RoleOwner:
		return !(level > 3)
	case user.RoleEmployee:
		return !(level > 2)
	default:
		return false
	}
}
