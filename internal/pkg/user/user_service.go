package user

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s *UserService) FindOne(id int) (*User, error) {
	var user User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &user, nil
}

func (s *UserService) FindByUsername(username string) (*User, error) {
	var user User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &user, nil
}

func (s *UserService) FindAll(query UserQuery) *pagination.Result[User] {
	result := pagination.New[User](query.Pagination)

	db := s.db.Model(&User{})

	if len(query.Role) > 0 {
		db.Where("role IN (?)", query.Role)
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *UserService) Create(data UserCreateDTO) (*User, error) {
	user := User{
		Name:     data.Name,
		Username: data.Username,
		Role:     data.Role,
		Status:   *data.Status,
	}

	user.SetPassword(data.Password)

	if err := s.db.Create(&user).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &user, nil
}

func (s *UserService) Update(id int, data UserUpdateDTO) (*User, error) {
	var user User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	user.Name = data.Name
	user.Username = data.Username
	user.Role = data.Role
	user.Status = *data.Status

	if data.Password != "" {
		user.SetPassword(data.Password)
	}

	if err := s.db.Save(&user).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &user, nil
}

func (s *UserService) Delete(id int) (*User, error) {
	var user User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&user).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &user, nil
}

func (s *UserService) Using(tx *gorm.DB) *UserService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *UserService) WithContext(ctx context.Context) *UserService {
	s.db = s.db.WithContext(ctx)

	return s
}
