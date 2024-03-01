package wage

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type WageController struct {
	*common.BaseController
	wage *WageService
}

func NewController(ctrl *common.BaseController, wage *WageService) *WageController {
	return &WageController{ctrl, wage}
}

// @Summary Get One Wage
// @Tags Wages
// @Accept json
// @Produce json
// @Param id path string true "Wage ID"
// @Success 200 {object} Wage{}
// @Security JWT
// @Router /api/wage/{id} [get]
func (ctrl *WageController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	wage, err := ctrl.wage.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(wage)
}

// @Summary Get All Wage
// @Tags Wages
// @Accept json
// @Produce json
// @Param query query WageQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Wage}
// @Security JWT
// @Router /api/wage [get]
func (ctrl *WageController) All(ctx *fiber.Ctx) error {
	var query WageQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.wage.WithContext(ctx.Context()).FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Wage
// @Tags Wages
// @Accept json
// @Produce json
// @Param request body WageDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Wage}
// @Security JWT
// @Router /api/wage [post]
func (ctrl *WageController) Create(ctx *fiber.Ctx) error {
	var data WageDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	wage, err := ctrl.wage.WithContext(ctx.Context()).Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Kasbon berhasil dibuat",
		Result:  wage,
	})
}

// @Summary Update Wage
// @Tags Wages
// @Accept json
// @Produce json
// @Param id path string true "Wage ID"
// @Param request body WageDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Wage}
// @Security JWT
// @Router /api/wage/{id} [put]
func (ctrl *WageController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data WageDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	wage, err := ctrl.wage.WithContext(ctx.Context()).Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Kasbon berhasil diubah",
		Result:  wage,
	})
}

// @Summary Delete Wage
// @Tags Wages
// @Accept json
// @Produce json
// @Param id path string true "Wage ID"
// @Success 200 {object} common.GeneralResponse{result=Wage}
// @Security JWT
// @Router /api/wage/{id} [delete]
func (ctrl *WageController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	wage, err := ctrl.wage.WithContext(ctx.Context()).Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Kasbon berhasil dihapus",
		Result:  wage,
	})
}

// @Summary Get Wages Summary
// @Tags Wages
// @Accept json
// @Produce json
// @Param query query WageSummaryQuery false "query"
// @Success 200 {object} []WageSummary
// @Security JWT
// @Router /api/wage/summary [get]
func (ctrl *WageController) GetSummary(ctx *fiber.Ctx) error {
	var query WageSummaryQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result, err := ctrl.wage.GetSummary(query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Cancel Wage
// @Tags Wages
// @Accept json
// @Produce json
// @Param id path string true "Wage ID"
// @Success 200 {object} common.BasicResponse{}
// @Security JWT
// @Router /api/wage/{id}/cancel [patch]
func (ctrl *WageController) Cancel(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	if err := ctrl.wage.SetStatus(id, "canceled"); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.BasicResponse{
		Message: "Kasbon berhasil dibatalkan",
	})
}
