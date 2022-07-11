package services_test

import (
	"testing"

	"github.com/gabrielporto8/banking-api/pkg/errs"
	"github.com/gabrielporto8/banking-api/pkg/models"
	"github.com/gabrielporto8/banking-api/pkg/repositories"
	"github.com/gabrielporto8/banking-api/pkg/services"
)

var (
	accounts = map[int64]*models.Account{
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
	repo = repositories.NewAccountRepository(accounts)
)

func TestGetAccountByID(t *testing.T) {
	service := services.NewAccountService(repo)

	t.Run("should return correct account", func(t *testing.T) {
		want := accounts[1].Cpf
		got, _ := service.GetAccountByID(1)

		if got.Cpf != want {
			t.Errorf("want %v, got %v", want, got.Cpf)
		}
	})

	t.Run("should return NotFoundError for not existent given ID", func(t *testing.T) {
		want := errs.NewNotFoundError(models.ErrAccountNotFound).Error()
		_, got := service.GetAccountByID(3)

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

func TestGetAccountByCPF(t *testing.T) {
	service := services.NewAccountService(repo)

	t.Run("should return correct account", func(t *testing.T) {
		want := accounts[1].Cpf
		got, _ := service.GetAccountByCPF(want)

		if got.Cpf != want {
			t.Errorf("want %v, got %v", want, got.Cpf)
		}
	})

	t.Run("should return correct account when CPF has special chars", func(t *testing.T) {
		want := accounts[1].Cpf
		got, _ := service.GetAccountByCPF("000.000.000-02")

		if got.Cpf != want {
			t.Errorf("want %v, got %v", want, got.Cpf)
		}
	})

	t.Run("should return NotFoundError for not existent given ID", func(t *testing.T) {
		want := errs.NewNotFoundError(models.ErrAccountNotFound).Error()
		_, got := service.GetAccountByCPF("00000000003")

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

func TestCreateAccount(t *testing.T) {
	repository := repositories.NewAccountRepository(make(map[int64]*models.Account))
	service := services.NewAccountService(repository)

	t.Run("should create account", func(t *testing.T) {
		account := models.Account{
			Name: "Teste 3",
			Cpf: "00000000003",
			Secret: "password123",
			Balance: 5,
		}

		got := service.CreateAccount(&account)

		if got != nil {
			t.Errorf("want nil, got %v", got.Error())
		}
	})

	service = services.NewAccountService(repo)

	t.Run("should return ValidationError when name is empty", func(t *testing.T) {
		account := models.Account{}

		want := errs.NewValidationError(models.ErrNameRequired).Error()
		got := service.CreateAccount(&account)

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("should return ValidationError when CPF is empty", func(t *testing.T) {
		account := models.Account{
			Name: "Teste 3",
		}

		want := errs.NewValidationError(models.ErrCpfRequired).Error()
		got := service.CreateAccount(&account)

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("should return ValidationError when CPF is invalid", func(t *testing.T) {
		account := models.Account{
			Name: "Teste 3",
			Cpf: "000",
		}

		want := errs.NewValidationError(models.ErrInvalidCpf).Error()
		got := service.CreateAccount(&account)

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("should return ValidationError when Secret is invalid", func(t *testing.T) {
		account := models.Account{
			Name: "Teste 3",
			Cpf: "00000000003",
			Secret: "pass",
		}

		want := errs.NewValidationError(models.ErrInvalidSecret).Error()
		got := service.CreateAccount(&account)

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("should return ValidationError when Balance is invalid", func(t *testing.T) {
		account := models.Account{
			Name: "Teste 3",
			Cpf: "00000000003",
			Secret: "password123",
			Balance: -50,
		}

		want := errs.NewValidationError(models.ErrInvalidBalance).Error()
		got := service.CreateAccount(&account)

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("should return ConflictError when CPF already exists", func(t *testing.T) {
		account := models.Account{
			Name: "Teste 3",
			Cpf: "00000000002",
			Secret: "password123",
			Balance: 50,
		}

		want := errs.NewConflictError(models.ErrAccountCPFAlreadyExists).Error()
		got := service.CreateAccount(&account)

		if got.Error() != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}

func TestGetBalance(t *testing.T) {
	service := services.NewAccountService(repo)

	t.Run("return correct balance", func(t *testing.T) {
		want := float64(20)
		got, _ := service.GetBalance(1)

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("return ErrNotFound when not existing account ID", func(t *testing.T) {
		want := errs.NewNotFoundError(models.ErrAccountNotFound).Error()
		_, got := service.GetBalance(3)

		if got.Error() != want {
			t.Errorf("want %v, got %v", want, got.Error())
		}
	})
}