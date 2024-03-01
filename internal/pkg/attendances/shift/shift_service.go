package shift

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type ShiftService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *ShiftService {
	return &ShiftService{db}
}

func (s *ShiftService) FindOne(id int) (*Shift, error) {
	var shift Shift
	if err := s.db.First(&shift, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &shift, nil
}

func (s *ShiftService) FindAll(query ShiftQuery) *pagination.Result[Shift] {
	result := pagination.New[Shift](query.Pagination)

	db := s.db.Model(&Shift{})
	if query.Company != 0 {
		db.Where("company_id = ?", query.Company)
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *ShiftService) Create(data ShiftDTO) (*Shift, error) {
	shift := Shift{
		Name:        data.Name,
		Description: data.Description,
		StartTime:   data.StartTime,
		EndTime:     data.EndTime,
		CompanyID:   data.Company,
	}

	if err := s.db.Create(&shift).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &shift, nil
}

func (s *ShiftService) Update(id int, data ShiftDTO) (*Shift, error) {
	var shift Shift
	if err := s.db.First(&shift, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	shift.Name = data.Name
	shift.Description = data.Description
	shift.StartTime = data.StartTime
	shift.EndTime = data.EndTime
	shift.CompanyID = data.Company

	if err := s.db.Save(&shift).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &shift, nil
}

func (s *ShiftService) Delete(id int) (*Shift, error) {
	var shift Shift
	if err := s.db.First(&shift, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&shift).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &shift, nil
}

func (s *ShiftService) Using(tx *gorm.DB) *ShiftService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *ShiftService) WithContext(ctx context.Context) *ShiftService {
	s.db = s.db.WithContext(ctx)

	return s
}
