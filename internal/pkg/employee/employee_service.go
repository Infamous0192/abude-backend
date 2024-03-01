package employee

import (
	"abude-backend/internal/pkg/user"
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeeService struct {
	db   *gorm.DB
	user *user.UserService
}

func NewService(db *gorm.DB, user *user.UserService) *EmployeeService {
	return &EmployeeService{db, user}
}

func (s *EmployeeService) FindOne(id int) (*Employee, error) {
	var employee Employee
	if err := s.db.Preload("User").First(&employee, id).Error; err != nil {
		return nil, exception.DB(err, "Perusahaan")
	}

	return &employee, nil
}

func (s *EmployeeService) FindAll(query EmployeeQuery) *pagination.Result[Employee] {
	result := pagination.New[Employee](query.Pagination)

	db := s.db.Model(&Employee{}).Preload("User")

	if query.Keyword != "" {
		db.Where("name LIKE ?", "%"+query.Keyword+"%")
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *EmployeeService) Create(data EmployeeCreateDTO) (*Employee, error) {
	employee := Employee{
		Name:        data.Name,
		Phonenumber: data.Phonenumber,
		Address:     data.Address,
		Status:      *data.Status,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		user, err := s.user.Using(tx).Create(user.UserCreateDTO{
			Name:     data.Name,
			Username: data.Username,
			Role:     user.RoleEmployee,
			Password: data.Password,
			Status:   data.Status,
		})

		if err != nil {
			return err
		}

		employee.UserID = user.ID

		if err := tx.Omit(clause.Associations).Create(&employee).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.DB(err)
	}

	return &employee, nil
}

func (s *EmployeeService) Update(id int, data EmployeeUpdateDTO) (*Employee, error) {
	employee, err := s.FindOne(id)
	if err != nil {
		return nil, exception.DB(err)
	}

	employee.Name = data.Name
	employee.Address = data.Address
	employee.Phonenumber = data.Phonenumber
	employee.Status = *data.Status

	err = s.db.Transaction(func(tx *gorm.DB) error {
		user, err := s.user.Using(tx).Update(int(employee.UserID), user.UserUpdateDTO{
			Name:     data.Name,
			Username: data.Username,
			Role:     user.RoleEmployee,
			Password: data.Password,
			Status:   data.Status,
		})

		if err != nil {
			return err
		}

		employee.User = user
		employee.UserID = user.ID

		if err := tx.Omit(clause.Associations).Save(&employee).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.DB(err)
	}

	return employee, nil
}

func (s *EmployeeService) Delete(id int) (*Employee, error) {
	var employee Employee
	if err := s.db.Preload("User").First(&employee, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&employee).Error; err != nil {
			return err
		}

		if err := tx.Delete(&employee.User).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, exception.DB(err)
	}

	return &employee, nil
}

func (s *EmployeeService) Using(tx *gorm.DB) *EmployeeService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *EmployeeService) WithContext(ctx context.Context) *EmployeeService {
	s.db = s.db.WithContext(ctx)

	return s
}
