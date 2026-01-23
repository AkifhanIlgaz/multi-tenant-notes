package service

import (
	"errors"
	"time"

	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/models"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/ports"
	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "istikrarli_hayal_hakikattir"
const JWT_EXPIRATION = time.Hour * 24

type AuthService struct {
	userRepo   ports.UserRepository
	tenantRepo ports.TenantRepository
}

func NewAuthService(userRepo ports.UserRepository, tenantRepo ports.TenantRepository) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		tenantRepo: tenantRepo,
	}
}

func (s *AuthService) Login(email, password string, tenantSlug string) (models.User, error) {
	tenant, err := s.tenantRepo.GetTenantBySlug(tenantSlug)
	if err != nil {
		return models.User{}, errors.New("invalid tenant")
	}

	user, err := s.userRepo.GetUserByEmailAndPassword(email, password, tenant.Id)
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}

	// normalde hashli password ile kontrol edilecek
	if user.Password != password {
		return models.User{}, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) GenerateToken(user models.User) (string, error) {
	claims := models.JWTClaims{
		UserID:   user.Id,
		TenantID: user.TenantId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRATION)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SECRET))
}

func (s *AuthService) ValidateToken(tokenString string) (models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return models.JWTClaims{}, errors.New("failed to parse token: " + err.Error())
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return models.JWTClaims{}, errors.New("invalid token claims: failed to parse claims structure")
	}

	if !token.Valid {
		return models.JWTClaims{}, errors.New("token is expired or invalid")
	}

	return *claims, nil

}
