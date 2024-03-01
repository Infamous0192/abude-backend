package inventory

import (
	"abude-backend/internal/common"

	"github.com/gofiber/fiber/v2"
)

type InventoryController struct {
	*common.BaseController
	inventory *InventoryService
}

func NewController(ctrl *common.BaseController, inventory *InventoryService) *InventoryController {
	return &InventoryController{ctrl, inventory}
}

// @Summary Get One Inventory
// @Tags Inventories
// @Accept json
// @Produce json
// @Param id path string true "Inventory ID"
// @Success 200 {object} Inventory{}
// @Security JWT
// @Router /api/inventory/{id} [get]
func (ctrl *InventoryController) One(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	inventory, err := ctrl.inventory.FindOne(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(inventory)
}

// @Summary Get All Inventory
// @Tags Inventories
// @Accept json
// @Produce json
// @Param query query InventoryQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Inventory}
// @Security JWT
// @Router /api/inventory [get]
func (ctrl *InventoryController) All(ctx *fiber.Ctx) error {
	var query InventoryQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.inventory.FindAll(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Add Inventory
// @Tags Inventories
// @Accept json
// @Produce json
// @Param request body InventoryDTO true "Request Body"
// @Success 201 {object} common.BasicResponse{}
// @Security JWT
// @Router /api/inventory [put]
func (ctrl *InventoryController) Add(ctx *fiber.Ctx) error {
	var data InventoryDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	err := ctrl.inventory.Save(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.BasicResponse{
		Message: "Stock berhasil diubah",
	})
}

// @Summary Delete Inventory
// @Tags Inventories
// @Accept json
// @Produce json
// @Param id path string true "Inventory ID"
// @Success 200 {object} common.GeneralResponse{result=Inventory}
// @Security JWT
// @Router /api/inventory/{id} [delete]
func (ctrl *InventoryController) Delete(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	inventory, err := ctrl.inventory.Delete(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(common.GeneralResponse{
		Message: "Inventaris berhasil dihapus",
		Result:  inventory,
	})
}

// @Summary Get Stocks
// @Tags Inventories
// @Accept json
// @Produce json
// @Param query query StockQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Stock}
// @Security JWT
// @Router /api/inventory/stock [get]
func (ctrl *InventoryController) GetStock(ctx *fiber.Ctx) error {
	var query StockQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.inventory.GetStock(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Get Stocks
// @Tags Inventories
// @Accept json
// @Produce json
// @Param query query StockSummaryQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]StockSummary}
// @Security JWT
// @Router /api/inventory/summary [get]
func (ctrl *InventoryController) GetStockSummary(ctx *fiber.Ctx) error {
	var query StockSummaryQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result, err := ctrl.inventory.GetStockSummary(query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Get Recapitulation
// @Tags Inventories
// @Accept json
// @Produce json
// @Param id path string true "Recap ID"
// @Success 200 {object} Recapitulation{}
// @Security JWT
// @Router /api/inventory/recapitulation/:id [get]
func (ctrl *InventoryController) GetRecap(ctx *fiber.Ctx) error {
	id, err := ctrl.Validation.ParamsInt(ctx)
	if err != nil {
		return err
	}

	recap, err := ctrl.inventory.GetRecap(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(recap)
}

// @Summary Get Recapitulations
// @Tags Inventories
// @Accept json
// @Produce json
// @Param query query RecapitulationQuery false "query"
// @Success 200 {object} common.PaginatedResponse{result=[]Recapitulation}
// @Security JWT
// @Router /api/inventory [get]
func (ctrl *InventoryController) GetRecaps(ctx *fiber.Ctx) error {
	var query RecapitulationQuery
	if err := ctrl.Validation.Query(&query, ctx); err != nil {
		return err
	}

	result := ctrl.inventory.GetRecaps(query)

	return ctx.Status(fiber.StatusOK).JSON(result)
}

// @Summary Create Recapitulation
// @Tags Inventories
// @Accept json
// @Produce json
// @Param request body RecapitulationDTO true "Request Body"
// @Success 201 {object} common.GeneralResponse{result=Inventory}
// @Security JWT
// @Router /api/inventory [post]
func (ctrl *InventoryController) CreateRecap(ctx *fiber.Ctx) error {
	var data RecapitulationDTO
	if err := ctrl.Validation.Body(&data, ctx); err != nil {
		return err
	}

	recap, err := ctrl.inventory.CreateRecap(data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(common.GeneralResponse{
		Message: "Rekapitulasi berhasil dibuat",
		Result:  recap,
	})
}
