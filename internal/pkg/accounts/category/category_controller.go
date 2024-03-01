package category

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	*common.BaseController
	category *CategoryService
}

func NewController(ctrl *common.BaseController, category *CategoryService) *CategoryController {
	return &CategoryController{ctrl, category}
}

// @Summary Get One Category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} Category{}
// @Security JWT
// @Router /api/category/{id} [get]
func (ctrl *CategoryController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	category, err := ctrl.category.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(category)
}

// @Summary Get All Category
// @Tags Categories
// @Accept json
// @Produce json
// @Param query query CategoryQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Category}
// @Security JWT
// @Router /api/category [get]
func (ctrl *CategoryController) All(ctx *fiber.Ctx) error {
	var query CategoryQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.category.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Category
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body CategoryDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Category}
// @Security JWT
// @Router /api/category [post]
func (ctrl *CategoryController) Create(ctx *fiber.Ctx) error {
	var data CategoryDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	category, err := ctrl.category.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Kategori berhasil dibuat",
		Result:  category,
	})
}

// @Summary Update Category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param request body CategoryDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Category}
// @Security JWT
// @Router /api/category/{id} [put]
func (ctrl *CategoryController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data CategoryDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	category, err := ctrl.category.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Kategori berhasil diubah",
		Result:  category,
	})
}

// @Summary Delete Category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} common.GeneralResponse{result=Category}
// @Security JWT
// @Router /api/category/{id} [delete]
func (ctrl *CategoryController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	category, err := ctrl.category.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Kategori berhasil dihapus",
		Result:  category,
	})
}
