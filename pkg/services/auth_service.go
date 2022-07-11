package services

import (
	"errors"

	"github.com/gabrielporto8/banking-api/pkg/errs"
	"github.com/gabrielporto8/banking-api/pkg/models"
)

var ErrUnauthorized = errors.New("Unauthorized.")

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
		return nil, errs.NewUnauthorizedError(ErrUnauthorized)
	}

	err = account.CheckPassword(authentication.Secret)
	if err != nil {
		return nil, errs.NewUnauthorizedError(ErrUnauthorized)
	}
	token, err := s.jwtService.GenerateJWT(account.Cpf)
	if err != nil {
		return nil, err
	}
	return token, nil
}