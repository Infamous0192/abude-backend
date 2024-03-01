package shift

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type ShiftController struct {
	*common.BaseController
	shift *ShiftService
}

func NewController(ctrl *common.BaseController, shift *ShiftService) *ShiftController {
	return &ShiftController{ctrl, shift}
}

// @Summary Get One Shift
// @Tags Shifts
// @Accept json
// @Produce json
// @Param id path string true "Shift ID"
// @Success 200 {object} Shift{}
// @Security JWT
// @Router /api/shift/{id} [get]
func (ctrl *ShiftController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	shift, err := ctrl.shift.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(shift)
}

// @Summary Get All Shift
// @Tags Shifts
// @Accept json
// @Produce json
// @Param query query ShiftQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Shift}
// @Security JWT
// @Router /api/shift [get]
func (ctrl *ShiftController) All(ctx *fiber.Ctx) error {
	var query ShiftQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.shift.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Shift
// @Tags Shifts
// @Accept json
// @Produce json
// @Param request body ShiftDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Shift}
// @Security JWT
// @Router /api/shift [post]
func (ctrl *ShiftController) Create(ctx *fiber.Ctx) error {
	var data ShiftDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	shift, err := ctrl.shift.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Shift berhasil dibuat",
		Result:  shift,
	})
}

// @Summary Update Shift
// @Tags Shifts
// @Accept json
// @Produce json
// @Param id path string true "Shift ID"
// @Param request body ShiftDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Shift}
// @Security JWT
// @Router /api/shift/{id} [put]
func (ctrl *ShiftController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data ShiftDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	shift, err := ctrl.shift.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Shift berhasil diubah",
		Result:  shift,
	})
}

// @Summary Delete Shift
// @Tags Shifts
// @Accept json
// @Produce json
// @Param id path string true "Shift ID"
// @Success 200 {object} common.GeneralResponse{result=Shift}
// @Security JWT
// @Router /api/shift/{id} [delete]
func (ctrl *ShiftController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	shift, err := ctrl.shift.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Shift berhasil dihapus",
		Result:  shift,
	})
}
