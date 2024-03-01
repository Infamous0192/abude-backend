package outlet

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/auth"
	"abude-backend/internal/pkg/user"

	"github.com/gofiber/fiber/v2"
)

type OutletController struct {
	*common.BaseController
	outlet *OutletService
}

func NewController(ctrl *common.BaseController, outlet *OutletService) *OutletController {
	return &OutletController{ctrl, outlet}
}

// @Summary Get One Outlet
// @Tags Outlets
// @Accept json
// @Produce json
// @Param id path string true "Outlet ID"
// @Success 200 {object} Outlet{}
// @Security JWT
// @Router /api/outlet/{id} [get]
func (ctrl *OutletController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	outlet, err := ctrl.outlet.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(outlet)
}

// @Summary Get All Outlet
// @Tags Outlets
// @Accept json
// @Produce json
// @Param query query OutletQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Outlet}
// @Security JWT
// @Router /api/outlet [get]
func (ctrl *OutletController) All(ctx *fiber.Ctx) error {
	var query OutletQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	creds := auth.GetCreds(ctx.Context())
	if creds.Role == user.RoleEmployee {
		query.Employee = creds.ID
	}

	result := ctrl.outlet.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Outlet
// @Tags Outlets
// @Accept json
// @Produce json
// @Param request body OutletDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Outlet}
// @Security JWT
// @Router /api/outlet [post]
func (ctrl *OutletController) Create(ctx *fiber.Ctx) error {
	var data OutletDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	outlet, err := ctrl.outlet.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Outlet berhasil dibuat",
		Result:  outlet,
	})
}

// @Summary Update Outlet
// @Tags Outlets
// @Accept json
// @Produce json
// @Param id path string true "Outlet ID"
// @Param request body OutletDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Outlet}
// @Security JWT
// @Router /api/outlet/{id} [put]
func (ctrl *OutletController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data OutletDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	outlet, err := ctrl.outlet.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Outlet berhasil diubah",
		Result:  outlet,
	})
}

// @Summary Delete Outlet
// @Tags Outlets
// @Accept json
// @Produce json
// @Param id path string true "Outlet ID"
// @Success 200 {object} common.GeneralResponse{result=Outlet}
// @Security JWT
// @Router /api/outlet/{id} [delete]
func (ctrl *OutletController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	outlet, err := ctrl.outlet.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Outlet berhasil dihapus",
		Result:  outlet,
	})
}

// @Summary Get Outlet Employee
// @Tags Outlets
// @Accept json
// @Produce json
// @Param outletId path string true "Outlet ID"
// @Param employeeId path string true "Employee ID"
// @Success 200 {object} OutletEmployee{}
// @Security JWT
// @Router /api/outlet/{outletId}/employee/{employeeId} [get]
func (ctrl *OutletController) GetEmployee(ctx *fiber.Ctx) error {
	outletId, err := ctrl.Validation.ParamsInt(ctx, "outletId")
	if err != nil {
		return err
	}

	employeeId, err := ctrl.Validation.ParamsInt(ctx, "employeeId")
	if err != nil {
		return err
	}

	outlet, err := ctrl.outlet.GetEmployee(outletId, employeeId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(outlet)
}

// @Summary Get Outlet Employees
// @Tags Outlets
// @Accept json
// @Produce json
// @Param outletId path string true "Outlet ID"
// @Param query query OutletEmployeeQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]OutletEmployee}
// @Security JWT
// @Router /api/outlet/{outletId}/employee [get]
func (ctrl *OutletController) GetEmployees(ctx *fiber.Ctx) error {
	outletId, err := ctrl.Validation.ParamsInt(ctx, "outletId")
	if err != nil {
		return err
	}

	var query OutletEmployeeQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	query.Outlet = outletId

	result := ctrl.outlet.GetEmployees(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Add Outlet Employee
// @Tags Outlets
// @Accept json
// @Produce json
// @Param outletId path string true "Outlet ID"
// @Param request body OutletEmployeeDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=OutletEmployee}
// @Security JWT
// @Router /api/outlet/{outletId}/employee [put]
func (ctrl *OutletController) AddEmployee(ctx *fiber.Ctx) error {
	outletId, err := ctrl.Validation.ParamsInt(ctx, "outletId")
	if err != nil {
		return err
	}

	employeeId, err := ctrl.Validation.ParamsInt(ctx, "employeeId")
	if err != nil {
		return err
	}

	data := OutletEmployeeDTO{
		Employee: employeeId,
		Outlet:   outletId,
	}

	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	employee, err := ctrl.outlet.AddEmployee(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Pegawai berhasil ditambahkan",
		Result:  employee,
	})
}

// @Summary Remove Outlet Employee
// @Tags Outlets
// @Accept json
// @Produce json
// @Param id path string true "Outlet ID"
// @Success 200 {object} common.GeneralResponse{result=Outlet}
// @Security JWT
// @Router /api/outlet/{outletId}/employee/{employeeId} [delete]
func (ctrl *OutletController) RemoveEmployee(ctx *fiber.Ctx) error {
	outletId, err := ctrl.Validation.ParamsInt(ctx, "outletId")
	if err != nil {
		return err
	}

	employeeId, err := ctrl.Validation.ParamsInt(ctx, "employeeId")
	if err != nil {
		return err
	}

	employee, err := ctrl.outlet.RemoveEmployee(outletId, employeeId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Pegawai berhasil dihapus",
		Result:  employee,
	})
}

// @Summary Get Outlet Count
// @Tags Outlets
// @Accept json
// @Produce json
// @Param query query OutletCountQuery false "query"
// @Success 200 {object} OutletCount
// @Security JWT
// @Router /api/outlet/count [get]
func (ctrl *OutletController) GetCount(ctx *fiber.Ctx) error {
	var query OutletCountQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.outlet.GetCount(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}
