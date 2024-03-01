package auth

import (
	"abude-backend/internal/common"
	"context"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	*common.BaseController
	service *AuthService
}

func NewController(ctrl *common.BaseController, service *AuthService) *AuthController {
	return &AuthController{
		BaseController: ctrl,
		service:        service,
	}
}

// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginDTO{} true "Request Body"
// @Success 200 {object} Authenticated{}
// @Security Tenant
// @Security JWT
// @Router /api/auth/login [post]
func (ctrl *AuthController) Login(ctx *fiber.Ctx) error {
	var data LoginDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	result, err := ctrl.service.WithContext(ctx).Login(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Verify Token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} Authenticated{}
// @Security Tenant
// @Security JWT
// @Router /api/auth/me [get]
func (ctrl *AuthController) Verify(ctx *fiber.Ctx) error {
	creds := GetCreds(ctx.Context())

	return ctx.Status(fiber.StatusOK).JSON(creds)
}

// Get User's Credentials
func GetCreds(ctx context.Context) *common.Creds {
	creds := ctx.Value(CredsKey).(common.Creds)

	return &creds
}
