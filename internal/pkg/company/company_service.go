package company

import (
	"abude-backend/internal/pkg/user"
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type CompanyService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *CompanyService {
	return &CompanyService{db}
}

func (s *CompanyService) FindOne(id int) (*Company, error) {
	var company Company
	if err := s.db.First(&company, id).Error; err != nil {
		return nil, exception.DB(err, "Perusahaan")
	}

	return &company, nil
}

func (s *CompanyService) FindAll(query CompanyQuery) *pagination.Result[Company] {
	result := pagination.New[Company](query.Pagination)

	db := s.db.Model(&Company{})
	if len(query.Owner) > 0 {
		db.Where("id IN (?)", s.db.
			Table("company_owners").
			Select("company_id").
			Where("user_id IN (?)", query.Owner))
	}

	if query.Keyword != "" {
		db.Where("name LIKE ?", "%"+query.Keyword+"%")
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *CompanyService) Create(data CompanyDTO) (*Company, error) {
	company := Company{
		Name:   data.Name,
		Region: data.Region,
	}

	if err := s.db.Create(&company).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &company, nil
}

func (s *CompanyService) Update(id int, data CompanyDTO) (*Company, error) {
	var company Company
	if err := s.db.First(&company, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	company.Name = data.Name
	company.Region = data.Region

	if err := s.db.Save(&company).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &company, nil
}

func (s *CompanyService) Delete(id int) (*Company, error) {
	var company Company
	if err := s.db.First(&company, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&company).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &company, nil
}

func (s *CompanyService) GetOwners(companyId int, query pagination.Pagination) *pagination.Result[user.User] {
	result := pagination.New[user.User](query)

	db := s.db.Model(&user.User{})

	db.Where("id IN (?)", s.db.Table("company_owners").Select("user_id").Where("company_id = ?", companyId))

	return result.Paginate(db)
}

func (s *CompanyService) AddOwner(companyId int, userId int) error {
	var owner user.User
	if err := s.db.First(&owner, userId).Error; err != nil {
		return exception.DB(err, "User")
	}

	if owner.Role != user.RoleOwner {
		return exception.Forbidden()
	}

	company, err := s.FindOne(companyId)
	if err != nil {
		return err
	}

	if err := s.db.Model(&company).Association("Owners").Append(owner); err != nil {
		return exception.DB(err)
	}

	return nil
}

func (s *CompanyService) RemoveOwner(companyId int, userId int) error {
	var owner user.User
	if err := s.db.First(&owner, userId).Error; err != nil {
		return exception.DB(err, "User")
	}

	company, err := s.FindOne(companyId)
	if err != nil {
		return err
	}

	if err := s.db.Model(&company).Association("Owners").Delete(owner); err != nil {
		return exception.DB(err)
	}

	return nil
}

func (s *CompanyService) Using(tx *gorm.DB) *CompanyService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *CompanyService) WithContext(ctx context.Context) *CompanyService {
	s.db = s.db.WithContext(ctx)

	return s
}
