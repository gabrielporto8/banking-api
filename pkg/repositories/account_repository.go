package repositories

import (
	"time"

	"github.com/gabrielporto8/banking-api/pkg/errs"
	"github.com/gabrielporto8/banking-api/pkg/models"
)

var (
	accounts map[int64]*models.Account = make(map[int64]*models.Account)
	accountsLastID int64 = 0
)

type AccountRepository struct {}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r AccountRepository) GetAccountByID(ID int64) (*models.Account, *errs.AppError) {
	accs, ok := accounts[ID]
	if !ok {
		return nil, errs.NewNotFoundError(models.ErrAccountNotFound)
	}
	return accs, nil
}

func (r AccountRepository) GetAccountByCPF(cpf string) (*models.Account, *errs.AppError) {
	for _, acc := range accounts {
		if acc.Cpf == cpf {
			return acc, nil
		}
	}

	return nil, errs.NewNotFoundError(models.ErrAccountNotFound)
}

func (r AccountRepository) GetAccounts() map[int64]*models.Account {
	return accounts
}

func (r AccountRepository) GetBalance(ID int64) (float64, *errs.AppError) {
	account, ok := accounts[ID]
	if !ok {
		return 0, errs.NewNotFoundError(models.ErrAccountNotFound)
	}

	return account.Balance, nil
}

func (r AccountRepository) SaveAccount(account *models.Account) *errs.AppError {
	account.ID = accountsLastID
	account.CreatedAt = time.Now()
	accounts[accountsLastID] = account
	accountsLastID++

	return nil
}

func (r AccountRepository) UpdateAccount(account *models.Account) (*models.Account, *errs.AppError) {
	_, ok := accounts[account.ID]
	if !ok {
		return nil, errs.NewNotFoundError(models.ErrAccountNotFound)
	}

	accounts[account.ID] = account
	return account, nil
}