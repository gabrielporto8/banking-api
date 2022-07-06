package repositories

import "github.com/gabrielporto8/stone-challenge/app/models"

type AuthRepository struct {}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r AuthRepository) CreateToken(token string) *models.Token {
	return &models.Token{
		Token: token,
	}
}