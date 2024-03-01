package user

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	*common.BaseController
	user *UserService
}

func NewController(ctrl *common.BaseController, user *UserService) *UserController {
	return &UserController{ctrl, user}
}

// @Summary Get One User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User{}
// @Security JWT
// @Router /api/user/{id} [get]
func (ctrl *UserController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	user, err := ctrl.user.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

// @Summary Get All User
// @Tags Users
// @Accept json
// @Produce json
// @Param query query UserQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]User}
// @Security JWT
// @Router /api/user [get]
func (ctrl *UserController) All(ctx *fiber.Ctx) error {
	var query UserQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.user.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create User
// @Tags Users
// @Accept json
// @Produce json
// @Param request body UserCreateDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=User}
// @Security JWT
// @Router /api/user [post]
func (ctrl *UserController) Create(ctx *fiber.Ctx) error {
	var data UserCreateDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	user, err := ctrl.user.Create(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "User berhasil dibuat",
		Result:  user,
	})
}

// @Summary Update User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body UserUpdateDTO true "Request Body"
// @Success 200 {object} common.GeneralResponse{result=User}
// @Security JWT
// @Router /api/user/{id} [put]
func (ctrl *UserController) Update(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	var data UserUpdateDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	user, err := ctrl.user.Update(id, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "User berhasil diubah",
		Result:  user,
	})
}

// @Summary Delete User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} common.GeneralResponse{result=User}
// @Security JWT
// @Router /api/user/{id} [delete]
func (ctrl *UserController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	user, err := ctrl.user.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "User berhasil dihapus",
		Result:  user,
	})
}
