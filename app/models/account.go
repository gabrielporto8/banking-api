package models

import (
	"time"
)

type Account struct {
	ID        int64    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
