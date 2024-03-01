package outlet

import (
	"abude-backend/internal/pkg/employee"
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type OutletService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *OutletService {
	return &OutletService{db}
}

func (s *OutletService) FindOne(id int) (*Outlet, error) {
	var outlet Outlet
	if err := s.db.Preload("Company").First(&outlet, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &outlet, nil
}

func (s *OutletService) FindAll(query OutletQuery) *pagination.Result[Outlet] {
	result := pagination.New[Outlet](query.Pagination)

	db := s.db.Model(&Outlet{}).Preload("Company")

	if query.Company != 0 {
		db.Where("company_id = ?", query.Company)
	}

	if query.Owner != 0 {
		db.Where("company_id IN (?)", s.db.Table("companies").
			Select("id").
			Joins("JOIN company_owners ON companies.id = company_owners.company_id").
			Where("user_id = ?", query.Owner))
	}

	if query.Keyword != "" {
		db.Where("name LIKE ?", "%"+query.Keyword+"%")
	}

	if query.Status != nil {
		db.Where("status = ?", query.Status)
	}

	if query.Employee != 0 {
		db.Where("id IN (?)", s.db.
			Select("outlet_employees.outlet_id").
			Table("outlet_employees").
			Joins("INNER JOIN employees ON employees.id=outlet_employees.employee_id").
			Where("employees.user_id = ?", query.Employee),
		)
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *OutletService) Create(data OutletDTO) (*Outlet, error) {
	outlet := Outlet{
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Status:    *data.Status,
		CompanyID: data.Company,
	}

	if err := s.db.Create(&outlet).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &outlet, nil
}

func (s *OutletService) Update(id int, data OutletDTO) (*Outlet, error) {
	var outlet Outlet
	if err := s.db.First(&outlet, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	outlet.Name = data.Name
	outlet.Address = data.Address
	outlet.Latitude = data.Latitude
	outlet.Longitude = data.Longitude
	outlet.Status = *data.Status
	outlet.CompanyID = data.Company

	if err := s.db.Save(&outlet).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &outlet, nil
}

func (s *OutletService) Delete(id int) (*Outlet, error) {
	var outlet Outlet
	if err := s.db.First(&outlet, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&outlet).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &outlet, nil
}

func (s *OutletService) GetCount(query OutletCountQuery) *OutletCount {
	var result OutletCount

	db := s.db.Model(&Outlet{})
	db.Select("COUNT(*) AS TotalCount, COUNT(CASE WHEN status = 1 THEN 1 END) AS ActiveCount, COUNT(CASE WHEN status = 0 THEN 1 END) AS InactiveCount")

	if query.Company != 0 {
		db.Where("company_id = ?", query.Company)
	}

	db.Find(&result)

	return &result
}

func (s *OutletService) GetEmployee(outletId int, employeeId int) (*OutletEmployee, error) {
	var outlet OutletEmployee
	if err := s.db.Preload("Employee").Preload("Outlet").
		Where("outlet_id = ? AND employee_id = ?", outletId, employeeId).
		First(&outlet).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &outlet, nil
}

func (s *OutletService) GetEmployees(query OutletEmployeeQuery) *pagination.Result[OutletEmployee] {
	result := pagination.New[OutletEmployee](query.Pagination)

	db := s.db.Model(&OutletEmployee{}).Preload("Employee").Preload("Employee.User")

	if query.Employee != 0 {
		db.Where("employee_id = ?", query.Employee)
	}

	if query.Type != "" {
		db.Where("type = ?", query.Type)
	}

	if query.Outlet != 0 {
		db.Where("outlet_id = ?", query.Outlet)
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *OutletService) AddEmployee(data OutletEmployeeDTO) (*OutletEmployee, error) {
	var employee OutletEmployee
	if err := s.db.Where(OutletEmployee{
		EmployeeID: uint(data.Employee),
		OutletID:   uint(data.Outlet),
	}).FirstOrCreate(&employee).Error; err != nil {
		return nil, exception.DB(err)
	}

	employee.Type = data.Type
	if err := s.db.Save(&employee).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &employee, nil
}

func (s *OutletService) RemoveEmployee(outletId int, employeeId int) (*OutletEmployee, error) {
	var employee OutletEmployee
	if err := s.db.Where("outlet_id = ? AND employee_id = ?", outletId, employeeId).First(&employee).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&employee).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &employee, nil
}

func (s *OutletService) GetAdmins(id int, query pagination.Pagination) *pagination.Result[employee.Employee] {
	result := pagination.New[employee.Employee](query)

	db := s.db.Model(&employee.Employee{}).Preload("User")

	db.Where("id IN (?)", s.db.Table("outlet_admins").Select("employee_id").Where("outlet_id = ?", id))

	return result.Paginate(db)
}

func (s *OutletService) AddAdmin(outletId int, employeeId int) error {
	outlet, err := s.FindOne(outletId)
	if err != nil {
		return err
	}

	var employee employee.Employee
	if err := s.db.First(&employee, employeeId).Error; err != nil {
		return exception.DB(err, "Admin")
	}

	if err := s.db.Model(&outlet).Association("Admins").Append(employee); err != nil {
		return exception.DB(err)
	}

	return nil
}

func (s *OutletService) RemoveAdmin(outletId int, employeeId int) error {
	outlet, err := s.FindOne(outletId)
	if err != nil {
		return err
	}

	var employee employee.Employee
	if err := s.db.First(&employee, employeeId).Error; err != nil {
		return exception.DB(err, "Admin")
	}

	if err := s.db.Model(&outlet).Association("Admins").Delete(employee); err != nil {
		return exception.DB(err)
	}

	return nil
}

func (s *OutletService) Using(tx *gorm.DB) *OutletService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *OutletService) WithContext(ctx context.Context) *OutletService {
	s.db = s.db.WithContext(ctx)

	return s
}
