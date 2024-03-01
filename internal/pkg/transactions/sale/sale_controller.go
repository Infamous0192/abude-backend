package sale

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

type SaleController struct {
	*common.BaseController
	sale *SaleService
}

func NewController(ctrl *common.BaseController, sale *SaleService) *SaleController {
	return &SaleController{ctrl, sale}
}

// @Summary Get One Sale
// @Tags Sales
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Success 200 {object} Sale{}
// @Security JWT
// @Router /api/sale/{id} [get]
func (ctrl *SaleController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	sale, err := ctrl.sale.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(sale)
}

// @Summary Get All Sale
// @Tags Sales
// @Accept json
// @Produce json
// @Param query query SaleQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Sale}
// @Security JWT
// @Router /api/sale [get]
func (ctrl *SaleController) All(ctx *fiber.Ctx) error {
	var query SaleQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.sale.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Sale
// @Tags Sales
// @Accept json
// @Produce json
// @Param request body SaleDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Sale}
// @Security JWT
// @Router /api/sale [post]
func (ctrl *SaleController) Create(ctx *fiber.Ctx) error {
	var data SaleDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	creds := auth.GetCreds(ctx.Context())
	data.User = creds.ID

	sale, err := ctrl.sale.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Penjualan berhasil dibuat",
		Result:  sale,
	})
}

// @Summary Update Sale
// @Tags Sales
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Param request body SaleDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Sale}
// @Security JWT
// @Router /api/sale/{id} [put]
func (ctrl *SaleController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data SaleDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	sale, err := ctrl.sale.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Penjualan berhasil diubah",
		Result:  sale,
	})
}

// @Summary Delete Sale
// @Tags Sales
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Success 200 {object} common.GeneralResponse{result=Sale}
// @Security JWT
// @Router /api/sale/{id} [delete]
func (ctrl *SaleController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	sale, err := ctrl.sale.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Penjualan berhasil dihapus",
		Result:  sale,
	})
}

// @Summary Get Sales Summary
// @Tags Sales
// @Accept json
// @Produce json
// @Param query query SaleSummaryQuery false "query"
// @Success 200 {object} []SaleSummary
// @Security JWT
// @Router /api/sale/summary [get]
func (ctrl *SaleController) GetSummary(ctx *fiber.Ctx) error {
	var query SaleSummaryQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result, err := ctrl.sale.GetSummary(query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Cancel Sale
// @Tags Sales
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Success 200 {object} common.BasicResponse{}
// @Security JWT
// @Router /api/sale/{id}/cancel [patch]
func (ctrl *SaleController) Cancel(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	if err := ctrl.sale.SetStatus(id, "canceled"); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.BasicResponse{
		Message: "Penjualan berhasil dibatalkan",
	})
}
