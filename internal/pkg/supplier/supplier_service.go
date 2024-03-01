package supplier

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type SupplierService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *SupplierService {
	return &SupplierService{db}
}

func (s *SupplierService) FindOne(id int) (*Supplier, error) {
	var supplier Supplier
	if err := s.db.Preload("Company").First(&supplier, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &supplier, nil
}

func (s *SupplierService) FindAll(query SupplierQuery) *pagination.Result[Supplier] {
	result := pagination.New[Supplier](query.Pagination)

	db := s.db.Model(&Supplier{}).Preload("Company")

	if query.Company != 0 {
		db.Where("company_id = ?", query.Company)
	}

	if query.Keyword != "" {
		db.Where("name LIKE ?", "%"+query.Keyword+"%")
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *SupplierService) Create(data SupplierDTO) (*Supplier, error) {
	supplier := Supplier{
		Name:        data.Name,
		Description: data.Description,
		CompanyID:   data.Company,
	}

	if err := s.db.Create(&supplier).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &supplier, nil
}

func (s *SupplierService) Update(id int, data SupplierDTO) (*Supplier, error) {
	var supplier Supplier
	if err := s.db.First(&supplier, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	supplier.Name = data.Name
	supplier.Description = data.Description
	supplier.CompanyID = data.Company

	if err := s.db.Save(&supplier).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &supplier, nil
}

func (s *SupplierService) Delete(id int) (*Supplier, error) {
	var supplier Supplier
	if err := s.db.First(&supplier, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&supplier).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &supplier, nil
}

func (s *SupplierService) Using(tx *gorm.DB) *SupplierService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *SupplierService) WithContext(ctx context.Context) *SupplierService {
	s.db = s.db.WithContext(ctx)

	return s
}
