package services

import (
	"github.com/gabrielporto8/banking-api/app/errs"
	"github.com/gabrielporto8/banking-api/app/models"
)

type AuthService struct {
	accountService *AccountService
	jwtService *JWTService
}

func NewAuthService(accountService *AccountService, jwtService *JWTService) *AuthService {
	return &AuthService{
		accountService: accountService,
		jwtService: jwtService,
	}
}

func (s AuthService) GenerateToken(authentication *models.Authentication) (*models.Token, *errs.AppError){
	account, err := s.accountService.GetAccountByCPF(authentication.Cpf)
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