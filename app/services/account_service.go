package services

import (
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

func (s AccountService) GetAccountByID(ID int64) (*models.Account, error) {
	return s.accountRepository.GetAccountByID(ID)
}

func (s AccountService) GetAccountByCPF(cpf string) (*models.Account, error) {
	return s.accountRepository.GetAccountByCPF(cpf)
}

func (s AccountService) GetAccounts() map[int64]*models.Account {
	return s.accountRepository.GetAccounts()
}

func (s AccountService) GetBalance(ID int64) (float64, error) {
	return s.accountRepository.GetBalance(ID)
}

func (s AccountService) CreateAccount(account *models.Account) error {
	if err := account.HashPassword(account.Secret); err != nil {
		return err
	}

	return s.accountRepository.SaveAccount(account)
}

func (s AccountService) UpdateAccount(account *models.Account) (*models.Account, error) {
	return s.accountRepository.UpdateAccount(account)
}