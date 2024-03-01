package handover

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type ProofService struct {
	db *gorm.DB
}

func NewProofService(db *gorm.DB) *ProofService {
	return &ProofService{db}
}

func (s *ProofService) FindOne(id int) (*Proof, error) {
	var proof Proof
	if err := s.db.Preload("Outlet").First(&proof, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &proof, nil
}

func (s *ProofService) FindAll(query ProofQuery) *pagination.Result[Proof] {
	result := pagination.New[Proof](query.Pagination)

	db := s.db.Model(&Proof{}).Preload("Outlet")

	if query.Outlet != 0 {
		db.Where("outlet_id = ?", query.Outlet)
	}

	if !query.StartDate.IsZero() {
		db.Where("date >= ?", query.StartDate)
	}

	if !query.EndDate.IsZero() {
		db.Where("date <= ?", query.EndDate)
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *ProofService) Create(data ProofDTO) (*Proof, error) {
	proof := Proof{
		Date:     data.Date,
		Evidence: data.Evidence,
		OutletID: data.Outlet,
	}

	if err := s.db.Create(&proof).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &proof, nil
}

func (s *ProofService) Update(id int, data ProofDTO) (*Proof, error) {
	var proof Proof
	if err := s.db.First(&proof, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	proof.Date = data.Date
	proof.Evidence = data.Evidence
	proof.OutletID = data.Outlet

	if err := s.db.Save(&proof).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &proof, nil
}

func (s *ProofService) Delete(id int) (*Proof, error) {
	var proof Proof
	if err := s.db.First(&proof, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&proof).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &proof, nil
}

func (s *ProofService) Using(tx *gorm.DB) *ProofService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *ProofService) WithContext(ctx context.Context) *ProofService {
	s.db = s.db.WithContext(ctx)

	return s
}
