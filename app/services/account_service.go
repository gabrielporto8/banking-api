package services

import (
	"time"

	"github.com/gabrielporto8/stone-challenge/app/models"
	"github.com/gabrielporto8/stone-challenge/app/repositories"
)

type AccountService struct {
	accountRepository *repositories.AccountRepository
}

func NewAccountService(accountRepository *repositories.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (s AccountService) GetAccountByID(ID int64) (*models.Account, bool) {
	return s.accountRepository.GetAccountByID(ID)
}

func (s AccountService) GetAccounts() map[int64]*models.Account {
	return s.accountRepository.GetAccounts()
}

func (s AccountService) GetBalance(ID int64) (float64, bool) {
	balance, err := s.accountRepository.GetBalance(ID)
	return balance, err
}

func (s AccountService) CreateAccount(account *models.Account) bool {
	if err := account.HashPassword(account.Secret); err != nil {
		return false
	}

	account.CreatedAt = time.Now()

	return s.accountRepository.SaveAccount(account)
}

func (s AccountService) UpdateAccount(account *models.Account) bool {
	return s.accountRepository.UpdateAccount(account)
}