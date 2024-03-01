package purchase

import (
	"abude-backend/internal/pkg/inventories/product"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"
	"time"

	"gorm.io/gorm"
)

type PurchaseService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *PurchaseService {
	return &PurchaseService{db}
}

func (s *PurchaseService) FindOne(id int) (*Purchase, error) {
	var purchase Purchase
	if err := s.db.Preload("User").Preload("Supplier").Preload("Items").Preload("Items.Product").First(&purchase, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &purchase, nil
}

func (s *PurchaseService) FindAll(query PurchaseQuery) *pagination.Result[Purchase] {
	result := pagination.New[Purchase](query.Pagination)

	db := s.db.Model(&Purchase{}).Preload("User").Preload("Supplier").Preload("Items").Preload("Items.Product")
	if query.Outlet != 0 {
		db.Where("id IN (?)", s.db.
			Table("outlet_purchases").
			Select("purchase_id").
			Where("outlet_id = ?", query.Outlet))
	}

	if len(query.Status) > 0 {
		db.Where("status IN (?)", query.Status)
	}

	if query.User != "" {
		db.Where("user_id = ?", query.User)
	}

	if !query.StartDate.IsZero() {
		db.Where("date >= ?", query.StartDate)
	}

	if !query.EndDate.IsZero() {
		db.Where("date <= ?", query.EndDate)
	}

	db.Order("date DESC")

	return result.Paginate(db)
}

func (s *PurchaseService) Create(data PurchaseDTO) (*Purchase, error) {
	purchase := Purchase{
		Note:       data.Note,
		Status:     StatusAccepted,
		Date:       time.Now(),
		UserID:     data.User,
		SupplierID: data.Supplier,
		Type:       data.Type,
	}

	if !data.Date.IsZero() {
		purchase.Date = data.Date
	}

	for _, item := range data.Items {
		purchaseItem := PurchaseItem{
			Quantity:  item.Quantity,
			ProductID: item.Product,
			Status:    false,
		}

		if item.Price != nil {
			purchaseItem.Price = *item.Price
		} else {
			var product product.Product
			if err := s.db.First(&product, item.Product).Error; err != nil {
				return nil, exception.DB(err)
			}
		}

		purchaseItem.Total = purchaseItem.Quantity * purchaseItem.Price
		purchase.Total += purchaseItem.Total
		purchase.Items = append(purchase.Items, purchaseItem)
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&purchase).Error; err != nil {
			return err
		}

		if data.Source == "outlet" {
			var outlet outlet.Outlet
			if err := s.db.First(&outlet, data.SourceID).Error; err != nil {
				return err
			}

			if err := tx.Create(&OutletPurchase{Purchase: &purchase, Outlet: &outlet}).Error; err != nil {
				return err
			}
		}

		// if data.Source == "warehouse" {
		// 	warehouse, err := Warehouse.FindOne(data.SourceID)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	if err := tx.Model(&warehouse).Association("Purchases").Append(&purchase); err != nil {
		// 		return err
		// 	}
		// }

		return nil
	})

	if err != nil {
		return nil, exception.DB(err)
	}

	return &purchase, nil
}

func (s *PurchaseService) Update(id int, data PurchaseDTO) (*Purchase, error) {
	var purchase Purchase
	if err := s.db.First(&purchase, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	purchase.Date = data.Date
	purchase.Note = data.Note

	if err := s.db.Save(&purchase).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &purchase, nil
}

func (s *PurchaseService) SetStatus(id int, status string) error {
	var purchase Purchase
	if err := s.db.First(&purchase, id).Error; err != nil {
		return exception.DB(err)
	}

	if purchase.Status == status {
		return exception.BadRequest("Status tidak berubah")
	}

	purchase.Status = status

	if err := s.db.Save(&purchase).Error; err != nil {
		return exception.DB(err)
	}

	return nil
}

func (s *PurchaseService) Delete(id int) (*Purchase, error) {
	var purchase Purchase
	if err := s.db.First(&purchase, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&purchase).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &purchase, nil
}

func (s *PurchaseService) GetSummary(query PurchaseSummaryQuery) ([]PurchaseSummary, error) {
	var summary []PurchaseSummary

	db := s.db.Model(&Purchase{})
	db.Select("products.id, products.name, SUM(purchase_items.quantity) AS quantity, SUM(purchase_items.total) AS total, DATE(purchases.date) AS date")
	db.Joins("INNER JOIN outlet_purchases ON purchases.id = outlet_purchases.purchase_id")
	db.Joins("RIGHT JOIN purchase_items ON purchases.id = purchase_items.purchase_id")
	db.Joins("INNER JOIN products ON products.id = purchase_items.product_id")
	db.Where("purchases.status != ?", "canceled")

	if query.StartDate != "" {
		db.Where("DATE(purchases.date) >= ?", query.StartDate)
	}

	if query.EndDate != "" {
		db.Where("DATE(purchases.date) <= ?", query.EndDate)
	}

	if len(query.Status) > 0 {
		db.Where("purchases.status IN (?)", query.Status)
	}

	if query.Outlet != 0 {
		db.Where("outlet_purchases.outlet_id = ?", query.Outlet)
	}

	db.Group("purchase_items.product_id, DATE(purchases.date)")
	db.Order("DATE(purchases.date) ASC")

	if err := db.Find(&summary).Error; err != nil {
		return nil, exception.DB(err)
	}

	return summary, nil
}

func (s *PurchaseService) Using(tx *gorm.DB) *PurchaseService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *PurchaseService) WithContext(ctx context.Context) *PurchaseService {
	s.db = s.db.WithContext(ctx)

	return s
}
