package product

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *ProductService {
	return &ProductService{db}
}

func (s *ProductService) FindOne(id int) (*Product, error) {
	var product Product
	if err := s.db.Preload("Company").Preload("Category").Preload("Ingredients").Preload("Ingredients.Ingredient").First(&product, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &product, nil
}

func (s *ProductService) FindAll(query ProductQuery) *pagination.Result[Product] {
	result := pagination.New[Product](query.Pagination)

	db := s.db.Model(&Product{}).Preload("Company").Preload("Category")
	if query.Company != 0 {
		db.Where("company_id = ?", query.Company)
	}

	if query.Type != "" {
		db.Where("type = ?", query.Type)
	}

	if query.Category != 0 {
		db.Where("category = ?", query.Category)
	}

	if query.Default != nil {
		db.Where("is_default = ?", query.Default)
	}

	if query.Keyword != "" {
		db.Where("name LIKE ?", "%"+query.Keyword+"%")
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *ProductService) Create(data ProductDTO) (*Product, error) {
	product := Product{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Unit:        data.Unit,
		CategoryID:  data.Category,
		CompanyID:   data.Company,
		IsDefault:   data.IsDefault,
		Type:        data.Type,
		Stock:       data.Stock,
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&product).Error; err != nil {
			return err
		}

		if len(data.Ingredients) == 0 {
			return nil
		}

		var ingredients []Ingredient
		for _, v := range data.Ingredients {
			ingredients = append(ingredients, Ingredient{
				Quantity:     v.Quantity,
				BaseID:       product.ID,
				IngredientID: v.Product,
			})
		}

		if err := tx.Create(&ingredients).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, exception.DB(err)
	}

	return s.FindOne(int(product.ID))
}

func (s *ProductService) Update(id int, data ProductDTO) (*Product, error) {
	var product Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	product.Name = data.Name
	product.Description = data.Description
	product.Price = data.Price
	product.Unit = data.Unit
	product.CategoryID = data.Category
	product.CompanyID = data.Company
	product.IsDefault = data.IsDefault
	product.Type = data.Type
	product.Stock = data.Stock

	var ingredients []Ingredient
	for _, v := range data.Ingredients {
		ingredients = append(ingredients, Ingredient{
			Quantity:     v.Quantity,
			BaseID:       product.ID,
			IngredientID: v.Product,
		})
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&product).Error; err != nil {
			return err
		}

		if err := tx.Where("base_id = ?", product.ID).Delete(&Ingredient{}).Error; err != nil {
			return err
		}

		if err := tx.Create(&ingredients).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, exception.DB(err)
	}

	return &product, nil
}

func (s *ProductService) Delete(id int) (*Product, error) {
	var product Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&product).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &product, nil
}

func (s *ProductService) Using(tx *gorm.DB) *ProductService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *ProductService) WithContext(ctx context.Context) *ProductService {
	s.db = s.db.WithContext(ctx)

	return s
}
