package account

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	*common.BaseController
	account *AccountService
}

func NewController(ctrl *common.BaseController, account *AccountService) *AccountController {
	return &AccountController{ctrl, account}
}

// @Summary Get One Account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} Account{}
// @Security JWT
// @Router /api/account/{id} [get]
func (ctrl *AccountController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	account, err := ctrl.account.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(account)
}

// @Summary Get All Account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param query query AccountQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Account}
// @Security JWT
// @Router /api/account [get]
func (ctrl *AccountController) All(ctx *fiber.Ctx) error {
	var query AccountQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.account.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param request body AccountDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Account}
// @Security JWT
// @Router /api/account [post]
func (ctrl *AccountController) Create(ctx *fiber.Ctx) error {
	var data AccountDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	account, err := ctrl.account.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Akun berhasil dibuat",
		Result:  account,
	})
}

// @Summary Update Account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param request body AccountDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Account}
// @Security JWT
// @Router /api/account/{id} [put]
func (ctrl *AccountController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data AccountDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	account, err := ctrl.account.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Akun berhasil diubah",
		Result:  account,
	})
}

// @Summary Delete Account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} common.GeneralResponse{result=Account}
// @Security JWT
// @Router /api/account/{id} [delete]
func (ctrl *AccountController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	account, err := ctrl.account.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Akun berhasil dihapus",
		Result:  account,
	})
}
