package employee

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	*common.BaseController
	employee *EmployeeService
}

func NewController(ctrl *common.BaseController, employee *EmployeeService) *EmployeeController {
	return &EmployeeController{ctrl, employee}
}

// @Summary Get One Employee
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} Employee{}
// @Security JWT
// @Router /api/employee/{id} [get]
func (ctrl *EmployeeController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	employee, err := ctrl.employee.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(employee)
}

// @Summary Get All Employee
// @Tags Employees
// @Accept json
// @Produce json
// @Param query query EmployeeQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Employee}
// @Security JWT
// @Router /api/employee [get]
func (ctrl *EmployeeController) All(ctx *fiber.Ctx) error {
	var query EmployeeQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.employee.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Employee
// @Tags Employees
// @Accept json
// @Produce json
// @Param request body EmployeeCreateDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Employee}
// @Security JWT
// @Router /api/employee [post]
func (ctrl *EmployeeController) Create(ctx *fiber.Ctx) error {
	var data EmployeeCreateDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	employee, err := ctrl.employee.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Pegawai berhasil dibuat",
		Result:  employee,
	})
}

// @Summary Update Employee
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param request body EmployeeUpdateDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Employee}
// @Security JWT
// @Router /api/employee/{id} [put]
func (ctrl *EmployeeController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data EmployeeUpdateDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	employee, err := ctrl.employee.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Pegawai berhasil diubah",
		Result:  employee,
	})
}

// @Summary Delete Employee
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} common.GeneralResponse{result=Employee}
// @Security JWT
// @Router /api/employee/{id} [delete]
func (ctrl *EmployeeController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	employee, err := ctrl.employee.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Pegawai berhasil dihapus",
		Result:  employee,
	})
}
