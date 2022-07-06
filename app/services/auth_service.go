package services

import (
	"github.com/gabrielporto8/stone-challenge/app/models"
	"github.com/gabrielporto8/stone-challenge/app/repositories"
)

type AuthService struct {
	authReposiroty *repositories.AuthRepository
	accountRepository *repositories.AccountRepository
	jwtService *JWTService
}

func NewAuthService(authRepository *repositories.AuthRepository, accountRepository *repositories.AccountRepository, jwtService *JWTService) *AuthService {
	return &AuthService{
		authReposiroty: authRepository,
		accountRepository: accountRepository,
		jwtService: jwtService,
	}
}

func (s AuthService) GenerateToken(authentication *models.Authentication) (*models.Token, error){
	account, err := s.accountRepository.GetAccountByCPF(authentication.Cpf)
	if err != nil {
		return nil, err
	}

	err = account.CheckPassword(authentication.Secret)
	if err != nil {
		return nil, err
	}
	token, err := s.jwtService.GenerateJWT(account.Cpf)
	if err != nil {
		return nil, err
	}
	return token, nil
}