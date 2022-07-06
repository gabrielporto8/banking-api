package services

import (
	"errors"
	"time"

	"github.com/gabrielporto8/stone-challenge/app/models"
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

func (s JWTService) GenerateJWT(cpf string) (token *models.Token, err error) {
	expirationTime := time.Now().Add(30 * time.Minute).Unix()
	claims := &JWTClaim{
		Cpf: cpf,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString(jwtSecretKey)

	return &models.Token{Token: tokenString}, err
}

func (s JWTService) ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return ErrParseClaims
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return ErrExpiredToken
	}
	return
}