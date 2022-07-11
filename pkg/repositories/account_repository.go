package repositories

import (
	"time"

	"github.com/gabrielporto8/banking-api/pkg/errs"
	"github.com/gabrielporto8/banking-api/pkg/models"
)

var accountsLastID int64 = 0

type AccountRepository struct {
	accounts map[int64]*models.Account
}

func NewAccountRepository(accounts map[int64]*models.Account) *AccountRepository {
	return &AccountRepository{
		accounts: accounts,
	}
}

func (r AccountRepository) GetAccountByID(ID int64) (*models.Account, *errs.AppError) {
	accs, ok := r.accounts[ID]
	if !ok {
		return nil, errs.NewNotFoundError(models.ErrAccountNotFound)
	}
	return accs, nil
}

func (r AccountRepository) GetAccountByCPF(cpf string) (*models.Account, *errs.AppError) {
	for _, acc := range r.accounts {
		if acc.Cpf == cpf {
			return acc, nil
		}
	}

	return nil, errs.NewNotFoundError(models.ErrAccountNotFound)
}

func (r AccountRepository) GetAccounts() map[int64]*models.Account {
	return r.accounts
}

func (r AccountRepository) GetBalance(ID int64) (float64, *errs.AppError) {
	account, ok := r.accounts[ID]
	if !ok {
		return 0, errs.NewNotFoundError(models.ErrAccountNotFound)
	}

	return account.Balance, nil
}

func (r AccountRepository) SaveAccount(account *models.Account) *errs.AppError {
	account.ID = accountsLastID
	account.CreatedAt = time.Now()
	r.accounts[accountsLastID] = account
	accountsLastID++

	return nil
}

func (r AccountRepository) UpdateAccount(account *models.Account) (*models.Account, *errs.AppError) {
	_, ok := r.accounts[account.ID]
	if !ok {
		return nil, errs.NewNotFoundError(models.ErrAccountNotFound)
	}

	r.accounts[account.ID] = account
	return account, nil
}