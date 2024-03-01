package expense

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type ExpenseController struct {
	*common.BaseController
	expense *ExpenseService
}

func NewController(ctrl *common.BaseController, expense *ExpenseService) *ExpenseController {
	return &ExpenseController{ctrl, expense}
}

// @Summary Get One Expense
// @Tags Expenses
// @Accept json
// @Produce json
// @Param id path string true "Expense ID"
// @Success 200 {object} Expense{}
// @Security JWT
// @Router /api/expense/{id} [get]
func (ctrl *ExpenseController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	expense, err := ctrl.expense.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(expense)
}

// @Summary Get All Expense
// @Tags Expenses
// @Accept json
// @Produce json
// @Param query query ExpenseQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Expense}
// @Security JWT
// @Router /api/expense [get]
func (ctrl *ExpenseController) All(ctx *fiber.Ctx) error {
	var query ExpenseQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.expense.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Expense
// @Tags Expenses
// @Accept json
// @Produce json
// @Param request body ExpenseDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Expense}
// @Security JWT
// @Router /api/expense [post]
func (ctrl *ExpenseController) Create(ctx *fiber.Ctx) error {
	var data ExpenseDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	expense, err := ctrl.expense.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Pengeluaran berhasil dibuat",
		Result:  expense,
	})
}

// @Summary Update Expense
// @Tags Expenses
// @Accept json
// @Produce json
// @Param id path string true "Expense ID"
// @Param request body ExpenseDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Expense}
// @Security JWT
// @Router /api/expense/{id} [put]
func (ctrl *ExpenseController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data ExpenseDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	expense, err := ctrl.expense.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Pengeluaran berhasil diubah",
		Result:  expense,
	})
}

// @Summary Delete Expense
// @Tags Expenses
// @Accept json
// @Produce json
// @Param id path string true "Expense ID"
// @Success 200 {object} common.GeneralResponse{result=Expense}
// @Security JWT
// @Router /api/expense/{id} [delete]
func (ctrl *ExpenseController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	expense, err := ctrl.expense.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Pengeluaran berhasil dihapus",
		Result:  expense,
	})
}

// @Summary Get Expenses Summary
// @Tags Expenses
// @Accept json
// @Produce json
// @Param query query ExpenseSummaryQuery false "query"
// @Success 200 {object} []ExpenseSummary
// @Security JWT
// @Router /api/expense/summary [get]
func (ctrl *ExpenseController) GetSummary(ctx *fiber.Ctx) error {
	var query ExpenseSummaryQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result, err := ctrl.expense.GetSummary(query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Cancel Expense
// @Tags Expenses
// @Accept json
// @Produce json
// @Param id path string true "Expense ID"
// @Success 200 {object} common.BasicResponse{}
// @Security JWT
// @Router /api/expense/{id}/cancel [patch]
func (ctrl *ExpenseController) Cancel(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	if err := ctrl.expense.SetStatus(id, "canceled"); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.BasicResponse{
		Message: "Pengeluaran berhasil dibatalkan",
	})
}
