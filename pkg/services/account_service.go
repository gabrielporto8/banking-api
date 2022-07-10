package services

import (
	"github.com/gabrielporto8/banking-api/pkg/errs"
	"github.com/gabrielporto8/banking-api/pkg/models"
	"github.com/gabrielporto8/banking-api/pkg/repositories"
	"github.com/gabrielporto8/banking-api/pkg/utils"
)

type AccountService struct {
	accountRepository *repositories.AccountRepository
}

func NewAccountService(accountRepository *repositories.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (s AccountService) GetAccountByID(ID int64) (*models.Account, *errs.AppError) {
	return s.accountRepository.GetAccountByID(ID)
}

func (s AccountService) GetAccountByCPF(cpf string) (*models.Account, *errs.AppError) {
	cpfSanitized := utils.OnlyNumbersString(cpf)
	return s.accountRepository.GetAccountByCPF(cpfSanitized)
}

func (s AccountService) GetAccounts() map[int64]*models.Account {
	return s.accountRepository.GetAccounts()
}

func (s AccountService) GetBalance(ID int64) (float64, *errs.AppError) {
	return s.accountRepository.GetBalance(ID)
}

func (s AccountService) CreateAccount(account *models.Account) *errs.AppError {
	err := account.Validate()
	if err != nil {
		return err
	}

	err = account.HashPassword(account.Secret)
	if err != nil {
		return err
	}

	cpfSanitized := utils.OnlyNumbersString(account.Cpf)

	acc, _ := s.GetAccountByCPF(cpfSanitized)
	if acc != nil {
		return errs.NewConflictError(models.ErrAccountCPFAlreadyExists)
	}

	account.Cpf = cpfSanitized

	return s.accountRepository.SaveAccount(account)
}

func (s AccountService) UpdateAccountBalanceFromTransfer(transfer *models.Transfer) *errs.AppError {
	accountOrigin, err := s.GetAccountByID(transfer.AccountOriginID)
	if err != nil {
		return err
	}

	accountDestination, err := s.GetAccountByID(transfer.AccountDestinationID)
	if err != nil {
		return err
	}
	
	accountOrigin.Balance, accountDestination.Balance = accountOrigin.Balance - transfer.Amount, accountDestination.Balance + transfer.Amount
	_, err = s.UpdateAccount(accountOrigin)
	if err != nil {
		return err
	}

	_, err = s.UpdateAccount(accountDestination)
	if err != nil {
		return err
	}
	
	return nil
}

func (s AccountService) UpdateAccount(account *models.Account) (*models.Account, *errs.AppError) {
	return s.accountRepository.UpdateAccount(account)
}