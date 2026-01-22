package service

import (
	"errors"
	"time"

	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/api/dto"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/entity"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/ports"
	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "your_jwt_secret"
const JWT_EXPIRATION = time.Hour * 24

type AuthService struct {
	userRepo ports.UserRepository
}

func NewAuthService(userRepo ports.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(email, password string) (entity.User, error) {
	user, err := s.userRepo.GetUserByEmailAndPassword(email, password)
	if err != nil {
		return entity.User{}, errors.New("invalid credentials")
	}

	// normalde hashli password ile kontrol edilecek
	if user.Password != password {
		return entity.User{}, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) GenerateToken(user entity.User) (string, error) {
	claims := dto.JWTClaims{
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

func (s *AuthService) ValidateToken(tokenString string) (*dto.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dto.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
