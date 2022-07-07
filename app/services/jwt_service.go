package services

import (
	"errors"
	"time"

	"github.com/gabrielporto8/banking-api/app/errs"
	"github.com/gabrielporto8/banking-api/app/models"
	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtSecretKey = []byte("supersecretjwtkey")

	ErrParseClaims = errors.New("couldn't parse claims")
	ErrExpiredToken = errors.New("token expired")
) 

type JWTService struct {}

type JWTClaim struct {
	Cpf string `json:"cpf"`
	jwt.StandardClaims
}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (s JWTService) GenerateJWT(cpf string) (*models.Token, *errs.AppError) {
	expirationTime := time.Now().Add(30 * time.Minute).Unix()
	claims := &JWTClaim{
		Cpf: cpf,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString(jwtSecretKey)
	if err != nil {
		return nil, errs.NewInternalError(err)
	}

	return &models.Token{Token: tokenString}, nil
}

func (s JWTService) ValidateToken(signedToken string) (*JWTClaim, *errs.AppError) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		},
	)
	if err != nil {
		return nil, errs.NewUnauthorizedError(err)
	}
	claims, ok := token.Claims.(*JWTClaim)
	
	if !ok {
		return nil, errs.NewUnauthorizedError(ErrParseClaims)
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errs.NewUnauthorizedError(ErrExpiredToken)
	}
	return claims, nil
}