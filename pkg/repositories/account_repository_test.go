package repositories_test

import (
	"testing"

	"github.com/gabrielporto8/banking-api/pkg/errs"
	"github.com/gabrielporto8/banking-api/pkg/models"
	"github.com/gabrielporto8/banking-api/pkg/repositories"
)

var accounts = map[int64]*models.Account{
	0: {
		ID: 0,
		Name: "Name Teste",
		Cpf: "00000000001",
		Secret: "password123",
		Balance: 10,
	},
	1: {
		ID: 1,
		Name: "Name Teste 2",
		Cpf: "00000000002",
		Secret: "password123",
		Balance: 20,
	},
}

func TestSaveAccount(t *testing.T) {
	repository := repositories.NewAccountRepository(make(map[int64]*models.Account))

	account := models.Account{
		Name: "Name Teste",
		Cpf: "00000000001",
		Secret: "password123",
		Balance: 10,
	}

	got := repository.SaveAccount(&account)

	if got != nil {
		t.Errorf("want nil, got '%v", got.Error())
	}
}

func TestGetAccounts(t *testing.T) {
	repository := repositories.NewAccountRepository(accounts)
	
	want := 2
	got := repository.GetAccounts()

	if len(got) != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetBalance(t *testing.T) {
	repository := repositories.NewAccountRepository(accounts)

	t.Run("should return account balance", func(t *testing.T) {
		want := float64(20)
		got, _ := repository.GetBalance(1)
	
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("should return NotFoundError for not existent given ID", func(t *testing.T) {
		want := errs.NewNotFoundError(models.ErrAccountNotFound).Error()
		_, got := repository.GetBalance(3)

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

func TestGetAccountByID(t *testing.T) {
	repository := repositories.NewAccountRepository(accounts)

	t.Run("should return correct account", func(t *testing.T) {
		want := accounts[0].Cpf
		got, _ := repository.GetAccountByID(0)

		if got.Cpf != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("should return NotFoundError for not existent given ID", func(t *testing.T) {
		want := errs.NewNotFoundError(models.ErrAccountNotFound).Error()
		_, got := repository.GetAccountByID(3)

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

func TestGetAccountByCPF(t *testing.T) {
	repository := repositories.NewAccountRepository(accounts)

	t.Run("should return correct account", func(t *testing.T) {
		want := accounts[1].Cpf
		got, _ := repository.GetAccountByCPF(want)

		if got.Cpf != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("should return NotFoundError for not existent given ID", func(t *testing.T) {
		want := errs.NewNotFoundError(models.ErrAccountNotFound).Error()
		_, got := repository.GetAccountByCPF("00000000003")

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

func TestUpdateAccount(t *testing.T) {
	repository := repositories.NewAccountRepository(accounts)

	t.Run("should update account", func(t *testing.T) {
		newCpf := "00000000003"
		account := accounts[1]
		account.Cpf = newCpf

		updatedAccount, _ := repository.UpdateAccount(account)
		want := newCpf
		got := updatedAccount.Cpf

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("should return NotFoundError for not existent account", func(t *testing.T) {
		account := models.Account{
			ID: 3,
			Name: "Teste 3",
			Cpf: "00000000003",
			Balance: 5,
		}

		want := errs.NewNotFoundError(models.ErrAccountNotFound).Error()
		_, got := repository.UpdateAccount(&account)
		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}