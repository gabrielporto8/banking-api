package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountOriginNotFound = errors.New("account origin not found")
	ErrAccountDestinationNotFound = errors.New("account destination not found")
	ErrInsufficientBalance = errors.New("origin account does not have sufficient balance")
)

type Account struct {
	ID        int64    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func (account *Account) HashPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(account.Secret), 14)
	if err != nil {
		return err
	}

	account.Secret = string(hashed)
	return nil
}

func (account *Account) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
