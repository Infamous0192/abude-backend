package category

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *CategoryService {
	return &CategoryService{db}
}

func (s *CategoryService) FindOne(id int) (*Category, error) {
	var category Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &category, nil
}

func (s *CategoryService) FindAll(query CategoryQuery) *pagination.Result[Category] {
	result := pagination.New[Category](query.Pagination)

	db := s.db.Model(&Category{})

	if query.Keyword != "" {
		keyword := "%" + query.Keyword + "%"
		db.Where("name LIKE ? OR code LIKE ?", keyword, keyword)
	}

	if query.Company != 0 {
		db.Where("company_id = ?", query.Company)
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *CategoryService) Create(data CategoryDTO) (*Category, error) {
	category := Category{
		Name:        data.Name,
		Description: data.Description,
		Code:        data.Code,
		Normal:      data.Normal,
		CompanyID:   &data.Company,
	}

	if err := s.db.Create(&category).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &category, nil
}

func (s *CategoryService) Update(id int, data CategoryDTO) (*Category, error) {
	var category Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	category.Name = data.Name
	category.Description = data.Description
	category.Code = data.Code
	category.Normal = data.Normal
	category.CompanyID = &data.Company

	if err := s.db.Save(&category).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &category, nil
}

func (s *CategoryService) Delete(id int) (*Category, error) {
	var category Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&category).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &category, nil
}

func (s *CategoryService) Using(tx *gorm.DB) *CategoryService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *CategoryService) WithContext(ctx context.Context) *CategoryService {
	s.db = s.db.WithContext(ctx)

	return s
}
