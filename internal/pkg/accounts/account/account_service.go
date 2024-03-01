package account

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *AccountService {
	return &AccountService{db}
}

func (s *AccountService) FindOne(id int) (*Account, error) {
	var account Account
	if err := s.db.First(&account, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &account, nil
}

func (s *AccountService) FindByCode(code string, companyId uint) (*Account, error) {
	var account Account
	if err := s.db.Where("code = ? AND company_id = ?", code, companyId).First(&account).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &account, nil
}

func (s *AccountService) FindAll(query AccountQuery) *pagination.Result[Account] {
	result := pagination.New[Account](query.Pagination)

	db := s.db.Model(&Account{})

	if query.Category != 0 {
		db.Where("category_id = ?", query.Category)
	}

	if query.Keyword != "" {
		keyword := "%" + query.Keyword + "%"
		db.Where("name LIKE ? OR code LIKE ?", keyword, keyword)
	}

	if query.Company != 0 {
		db.Where("company_id = ? OR company_id IS NULL", query.Company)
	}

	db.Order("code ASC")

	return result.Paginate(db)
}

func (s *AccountService) Create(data AccountDTO) (*Account, error) {
	account := Account{
		Name:        data.Name,
		Description: data.Description,
		Code:        data.Code,
		Normal:      data.Normal,
		CategoryID:  data.Category,
		CompanyID:   &data.Company,
	}

	if err := s.db.Create(&account).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &account, nil
}

func (s *AccountService) Update(id int, data AccountDTO) (*Account, error) {
	var account Account
	if err := s.db.First(&account, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	account.Name = data.Name
	account.Description = data.Description
	account.Code = data.Code
	account.Normal = data.Normal
	account.CategoryID = data.Category
	account.CompanyID = &data.Company

	if err := s.db.Save(&account).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &account, nil
}

func (s *AccountService) Delete(id int) (*Account, error) {
	var account Account
	if err := s.db.First(&account, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&account).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &account, nil
}

func (s *AccountService) Using(tx *gorm.DB) *AccountService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *AccountService) WithContext(ctx context.Context) *AccountService {
	s.db = s.db.WithContext(ctx)

	return s
}
