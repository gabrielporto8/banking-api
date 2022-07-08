package models

import (
	"errors"
	"time"

	"github.com/gabrielporto8/banking-api/app/errs"
)

var (
	ErrSameAccountID = errors.New("Origin and destination account IDs cannot be the same.")
	ErrInvalidAmount = errors.New("The amount entered is invalid.")
)

type Transfer struct {
	ID                   int64    `json:"id"`
	AccountOriginID      int64    `json:"account_origin_id"`
	AccountDestinationID int64    `json:"account_destination_id"`
	Amount               float64   `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

func (t Transfer) Validate() *errs.AppError {
	if t.AccountOriginID == t.AccountDestinationID {
		return errs.NewConflictError(ErrSameAccountID)
	}

	if t.Amount <= 0 {
		return errs.NewValidationError(ErrInvalidAmount)
	}

	return nil
}
