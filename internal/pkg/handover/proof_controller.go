package handover

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type ProofController struct {
	*common.BaseController
	proof *ProofService
}

func NewProofController(ctrl *common.BaseController, proof *ProofService) *ProofController {
	return &ProofController{ctrl, proof}
}

// @Summary Get One Proof
// @Tags Handovers
// @Accept json
// @Produce json
// @Param id path string true "Proof ID"
// @Success 200 {object} Proof{}
// @Security JWT
// @Router /api/handover/proof/{id} [get]
func (ctrl *ProofController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	proof, err := ctrl.proof.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(proof)
}

// @Summary Get All Proof
// @Tags Handovers
// @Accept json
// @Produce json
// @Param query query ProofQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Proof}
// @Security JWT
// @Router /api/handover/proof [get]
func (ctrl *ProofController) All(ctx *fiber.Ctx) error {
	var query ProofQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.proof.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Proof
// @Tags Handovers
// @Accept json
// @Produce json
// @Param request body ProofDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Proof}
// @Security JWT
// @Router /api/handover/proof [post]
func (ctrl *ProofController) Create(ctx *fiber.Ctx) error {
	var data ProofDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	proof, err := ctrl.proof.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Bukti berhasil dibuat",
		Result:  proof,
	})
}

// @Summary Update Proof
// @Tags Handovers
// @Accept json
// @Produce json
// @Param id path string true "Proof ID"
// @Param request body ProofDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Proof}
// @Security JWT
// @Router /api/handover/proof/{id} [put]
func (ctrl *ProofController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data ProofDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	proof, err := ctrl.proof.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Bukti berhasil diubah",
		Result:  proof,
	})
}

// @Summary Delete Proof
// @Tags Handovers
// @Accept json
// @Produce json
// @Param id path string true "Proof ID"
// @Success 200 {object} common.GeneralResponse{result=Proof}
// @Security JWT
// @Router /api/handover/proof/{id} [delete]
func (ctrl *ProofController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	proof, err := ctrl.proof.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Bukti berhasil dihapus",
		Result:  proof,
	})
}
