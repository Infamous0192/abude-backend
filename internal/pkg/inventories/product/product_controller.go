package product

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	*common.BaseController
	product *ProductService
}

func NewController(ctrl *common.BaseController, product *ProductService) *ProductController {
	return &ProductController{ctrl, product}
}

// @Summary Get One Product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} Product{}
// @Security JWT
// @Router /api/product/{id} [get]
func (ctrl *ProductController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	product, err := ctrl.product.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(product)
}

// @Summary Get All Product
// @Tags Products
// @Accept json
// @Produce json
// @Param query query ProductQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Product}
// @Security JWT
// @Router /api/product [get]
func (ctrl *ProductController) All(ctx *fiber.Ctx) error {
	var query ProductQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.product.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body ProductDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Product}
// @Security JWT
// @Router /api/product [post]
func (ctrl *ProductController) Create(ctx *fiber.Ctx) error {
	var data ProductDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	product, err := ctrl.product.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Produk berhasil dibuat",
		Result:  product,
	})
}

// @Summary Update Product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body ProductDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Product}
// @Security JWT
// @Router /api/product/{id} [put]
func (ctrl *ProductController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data ProductDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	product, err := ctrl.product.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Produk berhasil diubah",
		Result:  product,
	})
}

// @Summary Delete Product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} common.GeneralResponse{result=Product}
// @Security JWT
// @Router /api/product/{id} [delete]
func (ctrl *ProductController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	product, err := ctrl.product.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Produk berhasil dihapus",
		Result:  product,
	})
}
