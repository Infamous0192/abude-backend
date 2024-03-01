package supplier

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type SupplierController struct {
	*common.BaseController
	supplier *SupplierService
}

func NewController(ctrl *common.BaseController, supplier *SupplierService) *SupplierController {
	return &SupplierController{ctrl, supplier}
}

// @Summary Get One Supplier
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Success 200 {object} Supplier{}
// @Security JWT
// @Router /api/supplier/{id} [get]
func (ctrl *SupplierController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	supplier, err := ctrl.supplier.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(supplier)
}

// @Summary Get All Supplier
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param query query SupplierQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Supplier}
// @Security JWT
// @Router /api/supplier [get]
func (ctrl *SupplierController) All(ctx *fiber.Ctx) error {
	var query SupplierQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.supplier.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Supplier
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param request body SupplierDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Supplier}
// @Security JWT
// @Router /api/supplier [post]
func (ctrl *SupplierController) Create(ctx *fiber.Ctx) error {
	var data SupplierDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	supplier, err := ctrl.supplier.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Supplier berhasil dibuat",
		Result:  supplier,
	})
}

// @Summary Update Supplier
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Param request body SupplierDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Supplier}
// @Security JWT
// @Router /api/supplier/{id} [put]
func (ctrl *SupplierController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data SupplierDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	supplier, err := ctrl.supplier.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Supplier berhasil diubah",
		Result:  supplier,
	})
}

// @Summary Delete Supplier
// @Tags Suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Success 200 {object} common.GeneralResponse{result=Supplier}
// @Security JWT
// @Router /api/supplier/{id} [delete]
func (ctrl *SupplierController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	supplier, err := ctrl.supplier.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Supplier berhasil dihapus",
		Result:  supplier,
	})
}
