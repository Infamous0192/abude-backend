package turnover

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type TurnoverService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *TurnoverService {
	return &TurnoverService{db}
}

func (s *TurnoverService) FindOne(id int) (*Turnover, error) {
	var turnover Turnover
	if err := s.db.Preload("Outlet").First(&turnover, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &turnover, nil
}

func (s *TurnoverService) FindAll(query TurnoverQuery) *pagination.Result[Turnover] {
	result := pagination.New[Turnover](query.Pagination)

	db := s.db.Model(&Turnover{}).Preload("Outlet")
	if query.Outlet != 0 {
		db.Where("outlet_id = ?", query.Outlet)
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

func (s *TurnoverService) Create(data TurnoverDTO) (*Turnover, error) {
	turnover := Turnover{
		Date:     data.Date,
		Evidence: data.Evidence,
		OutletID: data.Outlet,
	}

	if s.isDuplicate(&turnover) {
		return nil, exception.Validation(map[string]string{
			"date": "Tanggal ini sudah dibuat",
		})
	}

	if err := s.db.Create(&turnover).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &turnover, nil
}

func (s *TurnoverService) Update(id int, data TurnoverDTO) (*Turnover, error) {
	var turnover Turnover
	if err := s.db.First(&turnover, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	turnover.Date = data.Date
	turnover.Evidence = data.Evidence
	turnover.OutletID = data.Outlet

	if s.isDuplicate(&turnover) {
		return nil, exception.Validation(map[string]string{
			"date": "Tanggal ini sudah dibuat",
		})
	}

	if err := s.db.Save(&turnover).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &turnover, nil
}

func (s *TurnoverService) Delete(id int) (*Turnover, error) {
	var turnover Turnover
	if err := s.db.First(&turnover, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&turnover).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &turnover, nil
}

func (s *TurnoverService) isDuplicate(turnover *Turnover) bool {
	var count int64
	s.db.Model(&turnover).
		Where("DATE(date) = ? AND outlet_id = ? AND id != ?", turnover.Date.Format("2006-01-02"), turnover.OutletID, turnover.ID).
		Count(&count)

	return count > 0
}

func (s *TurnoverService) Using(tx *gorm.DB) *TurnoverService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *TurnoverService) WithContext(ctx context.Context) *TurnoverService {
	s.db = s.db.WithContext(ctx)

	return s
}
