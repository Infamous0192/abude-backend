package inventory

import (
	"abude-backend/internal/pkg/transactions/purchase"
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type InventoryService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *InventoryService {
	return &InventoryService{db}
}

func (s *InventoryService) FindOne(id int) (*Inventory, error) {
	var inventory Inventory
	if err := s.db.First(&inventory, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &inventory, nil
}

func (s *InventoryService) FindAll(query InventoryQuery) *pagination.Result[Inventory] {
	result := pagination.New[Inventory](query.Pagination)

	db := s.db.Model(&Inventory{})

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *InventoryService) StockIn(data InventoryDTO) error {
	var sourceQuery *gorm.DB
	if data.Source == "outlet" {
		sourceQuery = s.db.Table("outlet_inventories").Select("inventory_id").Where("outlet_id = ?", data.SourceID)
	} else if data.Source == "warehouse" {
		sourceQuery = s.db.Table("warehouse_inventories").Select("inventory_id").Where("warehouse_id = ?", data.SourceID)
	}

	var inventory Inventory
	result := s.db.Where(Inventory{
		Date:      data.Date,
		Price:     data.Price,
		ProductID: data.Product,
	}).Where("id IN (?)", sourceQuery).
		Attrs(Inventory{StockIn: 0, StockOut: 0}).FirstOrCreate(&inventory)
	if result.Error != nil {
		return exception.DB(result.Error)
	}

	inventory.StockIn += data.Quantity

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&inventory).Error; err != nil {
			return err
		}

		if result.RowsAffected > 0 {
			if data.Source == "outlet" {
				if err := tx.Create(&OutletInventory{
					InventoryID: inventory.ID,
					OutletID:    data.SourceID,
				}).Error; err != nil {
					return err
				}
			}
		}

		return nil
	}); err != nil {
		return exception.DB(err)
	}

	if err := s.db.Save(&inventory).Error; err != nil {
		return exception.DB(err)
	}

	if result.RowsAffected > 0 {
		if data.Source == "outlet" {
			s.db.Create(&OutletInventory{
				InventoryID: inventory.ID,
				OutletID:    data.SourceID,
			})
		}
	}

	return nil
}

func (s *InventoryService) StockOut(data InventoryDTO) error {
	var sourceQuery *gorm.DB
	if data.Source == "outlet" {
		sourceQuery = s.db.Table("outlet_inventories").Select("inventory_id").Where("outlet_id = ?", data.SourceID)
	} else if data.Source == "warehouse" {
		sourceQuery = s.db.Table("warehouse_inventories").Select("inventory_id").Where("warehouse_id = ?", data.SourceID)
	}

	data.Quantity *= -1

	var available float64
	if err := s.db.
		Table("inventories").
		Select("SUM(stock_in) - SUM(stock_out) AS available").
		Where("product_id = ? AND id IN (?)", data.Product, sourceQuery).Row().Scan(&available); err != nil {
		return exception.DB(err)
	}

	if available < data.Quantity {
		return exception.BadRequest("Stock tidak cukup")
	}

	var inventories []Inventory
	if err := s.db.Table("inventories").
		Where("stock_in - stock_out > 0 AND product_id = ? AND id IN (?)", data.Product, sourceQuery).
		Order("DATE ASC").Find(&inventories).Error; err != nil {
		return exception.DB(err)
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		index := 0
		count := data.Quantity
		for count > 0 {
			v := inventories[index]
			diff := count - (v.StockIn - v.StockOut)
			index++

			if diff == 0 {
				if err := tx.Delete(&Inventory{}, v.ID).Error; err != nil {
					return err
				}

				count = 0
			} else if diff > 0 {
				if err := tx.Delete(&Inventory{}, v.ID).Error; err != nil {
					return err
				}

				count = diff
			} else {
				if err := tx.Model(&Inventory{}).Where("id = ?", v.ID).Update("stock_out", v.StockOut+count).Error; err != nil {
					return err
				}

				count = 0
			}
		}

		return nil
	}); err != nil {
		return exception.DB(err)
	}

	return nil
}

func (s *InventoryService) Save(data InventoryDTO) error {
	if data.Quantity > 0 {
		if err := s.StockIn(data); err != nil {
			return err
		}
	} else {
		if err := s.StockOut(data); err != nil {
			return err
		}
	}

	return nil
}

func (s *InventoryService) Update(id int, data InventoryUpdateDTO) (*Inventory, error) {
	var inventory Inventory
	if err := s.db.First(&inventory, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Save(&inventory).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &inventory, nil
}

func (s *InventoryService) Delete(id int) (*Inventory, error) {
	var inventory Inventory
	if err := s.db.First(&inventory, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&inventory).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &inventory, nil
}

func (s *InventoryService) GetStock(query StockQuery) *pagination.Result[Stock] {
	result := pagination.New[Stock](query.Pagination)

	db := s.db.Table("inventories").
		Select("products.*, categories.*, SUM(stock_in - stock_out) AS amount, SUM((stock_in - stock_out) * inventories.price) AS total_value, SUM((stock_in - stock_out) * inventories.price) / SUM(stock_in - stock_out) AS average_price").
		Joins("INNER JOIN products ON products.id=inventories.product_id").
		Joins("INNER JOIN categories ON categories.id=products.category_id").
		Group("product_id")

	if query.Product != 0 {
		db.Where("inventories.product_id = ?", query.Product)
	}

	if query.Outlet != 0 {
		db.Where("inventories.id IN (?)", s.db.Table("outlet_inventories").Select("inventory_id").Where("outlet_id = ?", query.Outlet))
	}

	return result.Paginate(db)
}

func (s *InventoryService) GetStockSummary(query StockSummaryQuery) ([]StockSummary, error) {
	var stocks []StockSummary

	purchaseQuery := s.db.Table("purchase_items").
		Select("products.id AS product_id, SUM(purchase_items.quantity) AS stock_in, SUM(purchase_items.total) AS value_in, 0 AS stock_out, 0 AS value_out").
		Joins("INNER JOIN products ON products.id = purchase_items.product_id").
		Group("purchase_items.product_id").Where("purchase_items.status = 0")

	saleQuery := s.db.Table("sale_items").
		Select("products.id AS product_id, 0 AS stock_in, 0 AS value_in, SUM(ingredients.quantity) AS stock_out, SUM(ingredients.quantity * products.price) AS value_out").
		Joins("INNER JOIN ingredients ON ingredients.base_id = sale_items.product_id").
		Joins("INNER JOIN products ON products.id = ingredients.ingredient_id").
		Group("ingredients.ingredient_id").Where("sale_items.status = 0")

	if query.Outlet != 0 {
		purchaseQuery.Where("purchase_items.purchase_id IN (?)", s.db.
			Table("outlet_purchases").
			Select("outlet_purchases.purchase_id").
			Where("outlet_purchases.outlet_id = ?", query.Outlet))
		saleQuery.Where("sale_items.sale_id IN (?)", s.db.
			Table("outlet_sales").
			Select("outlet_sales.sale_id").
			Where("outlet_sales.outlet_id = ?", query.Outlet))
	}

	inventoryQuery := s.db.Table("inventories").
		Select("product_id, SUM(stock_in - stock_out) AS available, SUM((stock_in - stock_out) * price) AS total_value").
		Group("product_id")

	if query.Outlet != 0 {
		inventoryQuery.Where("inventories.id IN (?)", s.db.
			Table("outlet_inventories").
			Select("outlet_inventories.inventory_id").
			Where("outlet_inventories.outlet_id = ?", query.Outlet),
		)
	}

	if err := s.db.Select("products.*, SUM(stock_in) AS stock_in, SUM(stock_out) AS stock_out, SUM(value_in) AS value_in, SUM(value_out) AS value_out, SUM(available) AS available, SUM(total_value) AS total_value").
		Table("(? UNION ?) AS s", saleQuery, purchaseQuery).
		Joins("INNER JOIN (?) AS i ON i.product_id = s.product_id", inventoryQuery).
		Joins("INNER JOIN products ON products.id=s.product_id").
		Group("s.product_id, i.product_id").
		Find(&stocks).Error; err != nil {
		return stocks, exception.DB(err)
	}

	return stocks, nil
}

func (s *InventoryService) GetRecaps(query RecapitulationQuery) *pagination.Result[Recapitulation] {
	result := pagination.New[Recapitulation](query.Pagination)

	db := s.db.Model(&Recapitulation{}).Preload("Items").Preload("Items.Product")

	if query.Outlet != 0 {
		db.Where("outlet_id = ?", query.Outlet)
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *InventoryService) GetRecap(id int) (*Recapitulation, error) {
	var recap Recapitulation
	if err := s.db.Preload("Items").Preload("Items.Product").First(&recap, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &recap, nil
}

func (s *InventoryService) CreateRecap(data RecapitulationDTO) (*Recapitulation, error) {
	recap := Recapitulation{
		Date:     data.Date,
		Notes:    data.Notes,
		Employee: data.Employee,
		OutletID: data.Outlet,
	}

	items, err := s.GetStockSummary(StockSummaryQuery{
		Outlet: int(data.Outlet),
	})
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.StockOut > (item.StockIn + item.Available) {
			return nil, exception.BadRequest(fmt.Sprintf("Stock '%s' tidak cukup", item.Product.Name))
		}

		recap.Items = append(recap.Items, RecapitulationItem{
			Available:  item.Available,
			TotalValue: item.TotalValue,
			StockIn:    item.StockIn,
			ValueIn:    item.ValueIn,
			ValueOut:   item.StockOut,
			StockOut:   item.StockOut,
			ProductID:  item.Product.ID,
		})
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&recap).Error; err != nil {
			return err
		}

		var purchases []purchase.PurchaseItem
		if err := tx.Where("status = 0 AND purchase_id IN (?)", s.db.
			Table("outlet_purchases").
			Select("purchase_id").
			Where("outlet_id = ?", data.Outlet),
		).Preload("Purchase").Find(&purchases).Error; err != nil {
			return err
		}

		if err := tx.Table("sale_items").
			Where("status = 0 AND sale_id IN (?)", s.db.Table("outlet_sales").Select("sale_id").Where("outlet_id = ?", data.Outlet)).
			Update("status", 1).Error; err != nil {
			return err
		}

		if err := tx.Table("purchase_items").
			Where("status = 0 AND purchase_id IN (?)", s.db.Table("outlet_purchases").Select("purchase_id").Where("outlet_id = ?", data.Outlet)).
			Update("status", 1).Error; err != nil {
			return err
		}

		service := NewService(tx)
		for _, v := range purchases {
			if err := service.StockIn(InventoryDTO{
				Source:   "outlet",
				SourceID: data.Outlet,
				Date:     datatypes.Date(v.Purchase.Date),
				Product:  v.ProductID,
				Price:    v.Price,
				Quantity: v.Quantity,
			}); err != nil {
				return err
			}
		}

		for _, v := range items {
			if v.StockOut == 0 {
				continue
			}

			if err := service.StockOut(InventoryDTO{
				Source:   "outlet",
				SourceID: data.Outlet,
				Date:     datatypes.Date(time.Now()),
				Product:  v.Product.ID,
				Price:    v.Product.Price,
				Quantity: v.StockOut,
			}); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, exception.DB(err)
	}

	return &recap, nil
}

func (s *InventoryService) Using(tx *gorm.DB) *InventoryService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *InventoryService) WithContext(ctx context.Context) *InventoryService {
	s.db = s.db.WithContext(ctx)

	return s
}
