package auth

import (
	"abude-backend/internal/common"
	"abude-backend/internal/config"
	"abude-backend/internal/pkg/user"
	"abude-backend/pkg/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	db     *gorm.DB
	config *config.JwtConfig
}

func NewService(db *gorm.DB, config *config.JwtConfig) *AuthService {
	return &AuthService{
		db:     db,
		config: config,
	}
}

func (s *AuthService) Login(data LoginDTO) (*Authenticated, error) {
	var user user.User
	if err := s.db.Where("username = ?", data.Username).First(&user).Error; err != nil {
		return nil, exception.Validation(map[string]string{
			"username": "Username tidak ditemukan",
		})
	}

	if !user.Status {
		return nil, exception.Forbidden()
	}

	if !user.ComparePassword(data.Password) {
		return nil, exception.Validation(map[string]string{
			"password": "Password salah",
		})
	}

	creds := common.Creds{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
		Status:   user.Status,
	}

	token, err := s.Hash(creds)
	if err != nil {
		return nil, err
	}

	return &Authenticated{
		Creds: creds,
		Token: token,
	}, nil
}

func (s *AuthService) Hash(creds common.Creds) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Creds: creds,
	})

	tokenString, err := token.SignedString([]byte(s.config.Secret))
	if err != nil {
		return "", exception.Http(fiber.StatusInternalServerError, err.Error())
	}

	return tokenString, nil
}

func (s *AuthService) Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.Secret), nil
	})

	if err != nil {
		return nil, exception.Http(fiber.StatusUnauthorized, "Token invalid or expired")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, exception.Http(fiber.StatusUnauthorized, "Token invalid or expired")
	}

	return claims, nil
}

func (s *AuthService) Using(tx *gorm.DB) *AuthService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx

	return s
}

func (s *AuthService) WithContext(ctx *fiber.Ctx) *AuthService {
	s.db = s.db.WithContext(ctx.Context())

	return s
}
