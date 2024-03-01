package sale

import (
	"abude-backend/internal/pkg/inventories/product"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SaleService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *SaleService {
	return &SaleService{db}
}

func (s *SaleService) FindOne(id int) (*Sale, error) {
	var awe []SaleTest
	s.db.Table("sale_items").Select("products.*, SUM(sale_items.total) AS total").Joins("INNER JOIN products ON products.id=sale_items.product_id").Group("product_id").Scan(&awe)

	fmt.Println(awe[0].Product.Name)

	var sale Sale
	if err := s.db.Preload("User").Preload("Items").Preload("Items.Product").First(&sale, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &sale, nil
}

func (s *SaleService) FindAll(query SaleQuery) *pagination.Result[Sale] {
	result := pagination.New[Sale](query.Pagination)

	db := s.db.Model(&Sale{}).Preload("User")
	if query.Outlet != 0 {
		db.Where("id IN (?)", s.db.
			Table("outlet_sales").
			Select("sale_id").
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

func (s *SaleService) Create(data SaleDTO) (*Sale, error) {
	sale := Sale{
		Customer: data.Customer,
		Note:     data.Note,
		Status:   StatusAccepted,
		Date:     time.Now(),
		UserID:   data.User,
	}

	if !data.Date.IsZero() {
		sale.Date = data.Date
	}

	if data.Status != nil {
		sale.Status = *data.Status
	}

	for _, item := range data.Items {
		saleItem := SaleItem{
			Quantity:  item.Quantity,
			ProductID: item.Product,
			Status:    false,
		}

		if item.Price != nil {
			saleItem.Price = *item.Price
		} else {
			var product product.Product
			if err := s.db.First(&product, item.Product).Error; err != nil {
				return nil, exception.DB(err)
			}
		}

		saleItem.Total = saleItem.Quantity * saleItem.Price
		sale.Total += saleItem.Total
		sale.Items = append(sale.Items, saleItem)
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&sale).Error; err != nil {
			return err
		}

		if data.Source == "outlet" {
			var outlet outlet.Outlet
			if err := s.db.First(&outlet, data.SourceID).Error; err != nil {
				return err
			}

			if err := tx.Create(&OutletSale{Sale: &sale, Outlet: &outlet}).Error; err != nil {
				return err
			}
		}

		// if data.Source == "warehouse" {
		// 	warehouse, err := Warehouse.FindOne(data.SourceID)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	if err := tx.Model(&warehouse).Association("Sales").Append(&sale); err != nil {
		// 		return err
		// 	}
		// }

		return nil
	})

	if err != nil {
		return nil, exception.DB(err)
	}

	return &sale, nil
}

func (s *SaleService) Update(id int, data SaleDTO) (*Sale, error) {
	var sale Sale
	if err := s.db.First(&sale, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	sale.Customer = data.Customer
	sale.Date = data.Date
	sale.Note = data.Note

	if err := s.db.Save(&sale).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &sale, nil
}

func (s *SaleService) SetStatus(id int, status string) error {
	var sale Sale
	if err := s.db.First(&sale, id).Error; err != nil {
		return exception.DB(err)
	}

	if sale.Status == status {
		return exception.BadRequest("Status tidak berubah")
	}

	sale.Status = status

	if err := s.db.Save(&sale).Error; err != nil {
		return exception.DB(err)
	}

	return nil
}

func (s *SaleService) Delete(id int) (*Sale, error) {
	var sale Sale
	if err := s.db.First(&sale, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&sale).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &sale, nil
}

func (s *SaleService) GetSummary(query SaleSummaryQuery) ([]SaleSummary, error) {
	var summary []SaleSummary

	db := s.db.Model(&Sale{})
	db.Select("products.id, products.name, SUM(sale_items.quantity) AS quantity, SUM(sale_items.total) AS total, DATE(sales.date) AS date")
	db.Joins("INNER JOIN outlet_sales ON sales.id = outlet_sales.sale_id")
	db.Joins("RIGHT JOIN sale_items ON sales.id = sale_items.sale_id")
	db.Joins("INNER JOIN products ON products.id = sale_items.product_id")
	db.Where("sales.status != ?", "canceled")

	if query.StartDate != "" {
		db.Where("DATE(sales.date) >= ?", query.StartDate)
	}

	if query.EndDate != "" {
		db.Where("DATE(sales.date) <= ?", query.EndDate)
	}

	if len(query.Status) > 0 {
		db.Where("sales.status IN (?)", query.Status)
	}

	if query.Outlet != 0 {
		db.Where("outlet_sales.outlet_id = ?", query.Outlet)
	}

	db.Group("sale_items.product_id, DATE(sales.date)")
	db.Order("DATE(sales.date) ASC")

	if err := db.Find(&summary).Error; err != nil {
		return nil, exception.DB(err)
	}

	return summary, nil
}

func (s *SaleService) Using(tx *gorm.DB) *SaleService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *SaleService) WithContext(ctx context.Context) *SaleService {
	s.db = s.db.WithContext(ctx)

	return s
}
