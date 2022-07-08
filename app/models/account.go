package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/gabrielporto8/banking-api/app/errs"
	"github.com/gabrielporto8/banking-api/app/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	CPFRules = regexp.MustCompile(`^\d{11}$`)

	ErrNameRequired = errors.New("The name field is required.")
	ErrCpfRequired = errors.New("The cpf field is required.")
	ErrInvalidCpf = errors.New("Invalid CPF format.")
	ErrInvalidSecret = errors.New("The secret must have at least 8 characters.")
	ErrInvalidBalance = errors.New("The balance must be a positive value.")

	ErrAccountNotFound = errors.New("Account not found.")
	ErrInsufficientBalance = errors.New("Origin account does not have sufficient balance.")
	ErrAccountCPFAlreadyExists = errors.New("Entered CPF already exists.")
)

type Account struct {
	ID        int64    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Account) HashPassword(password string) *errs.AppError {
	hashed, err := bcrypt.GenerateFromPassword([]byte(a.Secret), 14)
	if err != nil {
		return errs.NewInternalError(err)
	}

	a.Secret = string(hashed)
	return nil
}

func (a *Account) CheckPassword(providedPassword string) *errs.AppError {
	err := bcrypt.CompareHashAndPassword([]byte(a.Secret), []byte(providedPassword))
	if err != nil {
		return errs.NewUnauthorizedError(err)
	}
	return nil
}

func (a *Account) Validate() *errs.AppError {
	if a.Name == "" {
		return errs.NewValidationError(ErrNameRequired)
	}

	if a.Cpf == "" {
		return errs.NewValidationError(ErrCpfRequired)
	}

	if !CPFRules.MatchString(utils.OnlyNumbersString(a.Cpf)) {
		return errs.NewValidationError(ErrInvalidCpf)
	}

	if len(a.Secret) < 8 {
		return errs.NewValidationError(ErrInvalidSecret)
	}

	if a.Balance < 0 {
		return errs.NewValidationError(ErrInvalidBalance)
	}

	return nil
}

func (a *Account) ValidateTransferBalance(amount float64) *errs.AppError {
	if a.Balance < amount {
		return errs.NewValidationError(ErrInsufficientBalance)
	}

	return nil
}
