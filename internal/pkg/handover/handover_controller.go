package handover

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type HandoverController struct {
	*common.BaseController
	handover *HandoverService
}

func NewHandoverController(ctrl *common.BaseController, handover *HandoverService) *HandoverController {
	return &HandoverController{ctrl, handover}
}

// @Summary Get One Handover
// @Tags Handovers
// @Accept json
// @Produce json
// @Param id path string true "Handover ID"
// @Success 200 {object} Handover{}
// @Security JWT
// @Router /api/handover/{id} [get]
func (ctrl *HandoverController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	handover, err := ctrl.handover.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(handover)
}

// @Summary Get All Handover
// @Tags Handovers
// @Accept json
// @Produce json
// @Param query query HandoverQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Handover}
// @Security JWT
// @Router /api/handover [get]
func (ctrl *HandoverController) All(ctx *fiber.Ctx) error {
	var query HandoverQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.handover.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Handover
// @Tags Handovers
// @Accept json
// @Produce json
// @Param request body HandoverDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Handover}
// @Security JWT
// @Router /api/handover [post]
func (ctrl *HandoverController) Create(ctx *fiber.Ctx) error {
	var data HandoverDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	handover, err := ctrl.handover.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Serah Terima berhasil dibuat",
		Result:  handover,
	})
}

// @Summary Update Handover
// @Tags Handovers
// @Accept json
// @Produce json
// @Param id path string true "Handover ID"
// @Param request body HandoverDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Handover}
// @Security JWT
// @Router /api/handover/{id} [put]
func (ctrl *HandoverController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data HandoverDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	handover, err := ctrl.handover.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Serah Terima berhasil diubah",
		Result:  handover,
	})
}

// @Summary Delete Handover
// @Tags Handovers
// @Accept json
// @Produce json
// @Param id path string true "Handover ID"
// @Success 200 {object} common.GeneralResponse{result=Handover}
// @Security JWT
// @Router /api/handover/{id} [delete]
func (ctrl *HandoverController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	handover, err := ctrl.handover.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Serah Terima berhasil dihapus",
		Result:  handover,
	})
}
