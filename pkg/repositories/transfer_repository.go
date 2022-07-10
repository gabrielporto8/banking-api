package repositories

import (
	"github.com/gabrielporto8/banking-api/pkg/errs"
	"github.com/gabrielporto8/banking-api/pkg/models"
)

var (
	transfers map[int64]*models.Transfer = make(map[int64]*models.Transfer)
	transfersLastID int64 = 0
)

type TransferRepository struct {}

func NewTransferRepository() *TransferRepository {
	return &TransferRepository{}
}

func (r TransferRepository) GetTransfers() map[int64]*models.Transfer {
	return transfers
}

func (r TransferRepository) GetTransfersByOriginID(ID int64) []models.Transfer {
	var transfersFound []models.Transfer

	for _, transfer := range transfers {
		if transfer.AccountOriginID == ID {
			transfersFound = append(transfersFound, *transfer)
		}
	}

	return transfersFound
}

func (r TransferRepository) SaveTransfer(transfer *models.Transfer) *errs.AppError {
	transfer.ID = transfersLastID
	transfers[transfersLastID] = transfer
	transfersLastID++

	return nil
}