package repositories

import (
	"github.com/gabrielporto8/stone-challenge/app/models"
)

var (
	accounts map[int64]*models.Account = make(map[int64]*models.Account)
	accountsLastID int64 = 0
)

type AccountRepository struct {}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r AccountRepository) GetAccountByID(ID int64) (*models.Account, bool) {
	accs, ok := accounts[ID]
	return accs, ok
}

func (r AccountRepository) GetAccounts() map[int64]*models.Account {
	return accounts
}

func (r AccountRepository) GetBalance(ID int64) (float64, bool) {
	account, ok := accounts[ID]
	if !ok {
		return 0, ok
	}

	return account.Balance, ok
}

func (r AccountRepository) SaveAccount(account *models.Account) bool {
	account.ID = accountsLastID
	accounts[accountsLastID] = account
	accountsLastID++

	return true
}

func (r AccountRepository) UpdateAccount(account *models.Account) bool {
	_, ok := accounts[account.ID]
	if !ok {
		return false
	}

	accounts[account.ID] = account
	return true
}