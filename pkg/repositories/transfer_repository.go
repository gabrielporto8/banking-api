package repositories

import (
	"github.com/gabrielporto8/banking-api/pkg/errs"
	"github.com/gabrielporto8/banking-api/pkg/models"
)

var transfersLastID int64 = 0

type TransferRepository struct {
	transfers map[int64]*models.Transfer
}

func NewTransferRepository(transfers map[int64]*models.Transfer) *TransferRepository {
	return &TransferRepository{
		transfers: transfers,
	}
}

func (r TransferRepository) GetTransfers() map[int64]*models.Transfer {
	return r.transfers
}

func (r TransferRepository) GetTransfersByOriginID(ID int64) []models.Transfer {
	var transfersFound []models.Transfer

	for _, transfer := range r.transfers {
		if transfer.AccountOriginID == ID {
			transfersFound = append(transfersFound, *transfer)
		}
	}

	return transfersFound
}

func (r TransferRepository) SaveTransfer(transfer *models.Transfer) *errs.AppError {
	transfer.ID = transfersLastID
	r.transfers[transfersLastID] = transfer
	transfersLastID++

	return nil
}