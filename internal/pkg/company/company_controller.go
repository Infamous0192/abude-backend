package company

import (
	"abude-backend/internal/common"
	"abude-backend/pkg/pagination"

	"github.com/gofiber/fiber/v2"
)

type CompanyController struct {
	*common.BaseController
	company *CompanyService
}

func NewController(ctrl *common.BaseController, company *CompanyService) *CompanyController {
	return &CompanyController{ctrl, company}
}

// @Summary Get One Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} Company{}
// @Security JWT
// @Router /api/company/{id} [get]
func (ctrl *CompanyController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	company, err := ctrl.company.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(company)
}

// @Summary Get All Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param query query CompanyQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Company}
// @Security JWT
// @Router /api/company [get]
func (ctrl *CompanyController) All(ctx *fiber.Ctx) error {
	var query CompanyQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.company.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param request body CompanyDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Company}
// @Security JWT
// @Router /api/company [post]
func (ctrl *CompanyController) Create(ctx *fiber.Ctx) error {
	var data CompanyDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	company, err := ctrl.company.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Perusahaan berhasil dibuat",
		Result:  company,
	})
}

// @Summary Update Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Param request body CompanyDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=Company}
// @Security JWT
// @Router /api/company/{id} [put]
func (ctrl *CompanyController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data CompanyDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	company, err := ctrl.company.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Perusahaan berhasil diubah",
		Result:  company,
	})
}

// @Summary Delete Company
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} common.GeneralResponse{result=Company}
// @Security JWT
// @Router /api/company/{id} [delete]
func (ctrl *CompanyController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	company, err := ctrl.company.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Perusahaan berhasil dihapus",
		Result:  company,
	})
}

// @Summary Get Company Owners
// @Tags Companies
// @Accept json
// @Produce json
// @Param query query pagination.Pagination false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]user.User}
// @Security JWT
// @Router /api/company/:id/owner [get]
func (ctrl *CompanyController) GetOwners(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var query pagination.Pagination
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.company.GetOwners(id, query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Add Company Owner
// @Tags Companies
// @Accept json
// @Produce json
// @Param companyId path string true "Company ID"
// @Param userId path string true "User ID"
// @Success 200 {object} common.BasicResponse{}
// @Security JWT
// @Router /api/company/{companyId}/owner/{userId} [put]
func (ctrl *CompanyController) AddOwner(ctx *fiber.Ctx) error {
	userId, err := ctrl.Validation.ParamsInt(ctx, "userId")
	if err != nil {
		return err
	}

	companyId, err := ctrl.Validation.ParamsInt(ctx, "companyId")
	if err != nil {
		return err
	}

	if err := ctrl.company.AddOwner(companyId, userId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.BasicResponse{
		Message: "Owner berhasil ditambahkan",
	})
}

// @Summary Remove Company Owner
// @Tags Companies
// @Accept json
// @Produce json
// @Param companyId path string true "Company ID"
// @Param userId path string true "User ID"
// @Success 200 {object} common.BasicResponse{}
// @Security JWT
// @Router /api/company/{companyId}/owner/{userId} [delete]
func (ctrl *CompanyController) RemoveOwner(ctx *fiber.Ctx) error {
	userId, err := ctrl.Validation.ParamsInt(ctx, "userId")
	if err != nil {
		return err
	}

	companyId, err := ctrl.Validation.ParamsInt(ctx, "companyId")
	if err != nil {
		return err
	}

	if err := ctrl.company.AddOwner(companyId, userId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.BasicResponse{
		Message: "Owner berhasil dihapus",
	})
}
