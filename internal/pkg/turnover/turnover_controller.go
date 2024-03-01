package turnover

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type TurnoverController struct {
	*common.BaseController
	turnover *TurnoverService
}

func NewController(ctrl *common.BaseController, turnover *TurnoverService) *TurnoverController {
	return &TurnoverController{ctrl, turnover}
}

// @Summary Get One Turnover
// @Tags Turnovers
// @Accept json
// @Produce json
// @Param id path string true "Turnover ID"
// @Success 200 {object} Turnover{}
// @Security JWT
// @Router /api/turnover/{id} [get]
func (ctrl *TurnoverController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	turnover, err := ctrl.turnover.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(turnover)
}

// @Summary Get All Turnover
// @Tags Turnovers
// @Accept json
// @Produce json
// @Param query query TurnoverQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Turnover}
// @Security JWT
// @Router /api/turnover [get]
func (ctrl *TurnoverController) All(ctx *fiber.Ctx) error {
	var query TurnoverQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.turnover.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Turnover
// @Tags Turnovers
// @Accept json
// @Produce json
// @Param request body TurnoverDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Turnover}
// @Security JWT
// @Router /api/turnover [post]
func (ctrl *TurnoverController) Create(ctx *fiber.Ctx) error {
	var data TurnoverDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	turnover, err := ctrl.turnover.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Bukti berhasil dibuat",
		Result:  turnover,
	})
}

// @Summary Update Turnover
// @Tags Turnovers
// @Accept json
// @Produce json
// @Param id path string true "Turnover ID"
// @Param request body TurnoverDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Turnover}
// @Security JWT
// @Router /api/turnover/{id} [put]
func (ctrl *TurnoverController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data TurnoverDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	turnover, err := ctrl.turnover.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Bukti berhasil diubah",
		Result:  turnover,
	})
}

// @Summary Delete Turnover
// @Tags Turnovers
// @Accept json
// @Produce json
// @Param id path string true "Turnover ID"
// @Success 200 {object} common.GeneralResponse{result=Turnover}
// @Security JWT
// @Router /api/turnover/{id} [delete]
func (ctrl *TurnoverController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	turnover, err := ctrl.turnover.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Bukti berhasil dihapus",
		Result:  turnover,
	})
}
