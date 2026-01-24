package service

import (
	"context"
	"errors"
	"time"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/models"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/ports"
	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "bir_ilkbahar_sabahi"
const JWT_EXPIRATION = time.Hour * 24
const SchemaKey string = "schema"

type AuthService struct {
	userRepo ports.UserRepository
}

func NewAuthService(userRepo ports.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (models.User, error) {
	user, err := s.userRepo.GetUserByEmailAndPassword(ctx, email, password)
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}

	// normalde hashli password ile kontrol edilecek
	if user.Password != password {
		return models.User{}, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) GenerateToken(ctx context.Context, user models.User) (string, error) {
	schema, err := GetSchema(ctx)
	if err != nil {
		return "", err
	}

	claims := models.JWTClaims{
		UserID:     user.Id,
		TenantSlug: schema,
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

func GetSchema(ctx context.Context) (string, error) {
	schema, ok := ctx.Value(SchemaKey).(string)
	if !ok || schema == "" {
		return "", errors.New("schema not found in context")
	}
	return schema, nil
}
