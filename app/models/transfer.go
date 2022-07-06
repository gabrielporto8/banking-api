package models

import (
	"errors"
	"time"
)

var (
	ErrSameAccountID = errors.New("origin and destination account IDs cannot be the same")
	ErrInvalidAmount = errors.New("the amount entered is invalid")
)

type Transfer struct {
	ID                   int64    `json:"id"`
	AccountOriginID      int64    `json:"account_origin_id"`
	AccountDestinationID int64    `json:"account_destination_id"`
	Amount               float64   `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}
