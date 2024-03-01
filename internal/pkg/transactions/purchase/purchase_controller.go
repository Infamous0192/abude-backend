package purchase

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

type PurchaseController struct {
	*common.BaseController
	purchase *PurchaseService
}

func NewController(ctrl *common.BaseController, purchase *PurchaseService) *PurchaseController {
	return &PurchaseController{ctrl, purchase}
}

// @Summary Get One Purchase
// @Tags Purchases
// @Accept json
// @Produce json
// @Param id path string true "Purchase ID"
// @Success 200 {object} Purchase{}
// @Security JWT
// @Router /api/purchase/{id} [get]
func (ctrl *PurchaseController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	purchase, err := ctrl.purchase.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(purchase)
}

// @Summary Get All Purchase
// @Tags Purchases
// @Accept json
// @Produce json
// @Param query query PurchaseQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Purchase}
// @Security JWT
// @Router /api/purchase [get]
func (ctrl *PurchaseController) All(ctx *fiber.Ctx) error {
	var query PurchaseQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.purchase.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Purchase
// @Tags Purchases
// @Accept json
// @Produce json
// @Param request body PurchaseDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Purchase}
// @Security JWT
// @Router /api/purchase [post]
func (ctrl *PurchaseController) Create(ctx *fiber.Ctx) error {
	var data PurchaseDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	creds := auth.GetCreds(ctx.Context())
	data.User = creds.ID

	purchase, err := ctrl.purchase.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Penjualan berhasil dibuat",
		Result:  purchase,
	})
}

// @Summary Update Purchase
// @Tags Purchases
// @Accept json
// @Produce json
// @Param id path string true "Purchase ID"
// @Param request body PurchaseDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Purchase}
// @Security JWT
// @Router /api/purchase/{id} [put]
func (ctrl *PurchaseController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data PurchaseDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	purchase, err := ctrl.purchase.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Penjualan berhasil diubah",
		Result:  purchase,
	})
}

// @Summary Delete Purchase
// @Tags Purchases
// @Accept json
// @Produce json
// @Param id path string true "Purchase ID"
// @Success 200 {object} common.GeneralResponse{result=Purchase}
// @Security JWT
// @Router /api/purchase/{id} [delete]
func (ctrl *PurchaseController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	purchase, err := ctrl.purchase.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Penjualan berhasil dihapus",
		Result:  purchase,
	})
}

// @Summary Get Purchases Summary
// @Tags Purchases
// @Accept json
// @Produce json
// @Param query query PurchaseSummaryQuery false "query"
// @Success 200 {object} []PurchaseSummary
// @Security JWT
// @Router /api/purchase/summary [get]
func (ctrl *PurchaseController) GetSummary(ctx *fiber.Ctx) error {
	var query PurchaseSummaryQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result, err := ctrl.purchase.GetSummary(query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Cancel Purchase
// @Tags Purchases
// @Accept json
// @Produce json
// @Param id path string true "Purchase ID"
// @Success 200 {object} common.BasicResponse{}
// @Security JWT
// @Router /api/purchase/{id}/cancel [patch]
func (ctrl *PurchaseController) Cancel(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	if err := ctrl.purchase.SetStatus(id, "canceled"); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.BasicResponse{
		Message: "Penjualan berhasil dibatalkan",
	})
}
