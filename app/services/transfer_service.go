package services

import (
	"time"

	"github.com/gabrielporto8/banking-api/app/models"
	"github.com/gabrielporto8/banking-api/app/repositories"
)

type TransferService struct {
	transferRepository *repositories.TransferRepository
	accountService *AccountService
}

func NewTransferService(transferRepository *repositories.TransferRepository, accountService *AccountService) *TransferService {
	return &TransferService{
		transferRepository: transferRepository,
		accountService: accountService,
	}
}

func (s TransferService) GetTransfers() map[int64]*models.Transfer {
	return s.transferRepository.GetTransfers()
}

func (s TransferService) GetTransfersByOriginID(ID int64) []models.Transfer {
	return s.transferRepository.GetTransfersByOriginID(ID)
}

func (s TransferService) CreateTransfer(transfer *models.Transfer) error {
	if transfer.AccountOriginID == transfer.AccountDestinationID {
		return models.ErrSameAccountID
	}

	if transfer.Amount <= 0 {
		return models.ErrInvalidAmount
	}

	accountOrigin, err := s.accountService.GetAccountByID(transfer.AccountOriginID)
	if err != nil {
		return err
	}

	accountDestination, err := s.accountService.GetAccountByID(transfer.AccountDestinationID)
	if err != nil {
		return err
	}

	if accountOrigin.Balance < transfer.Amount {
		return models.ErrInsufficientBalance
	}

	transfer.CreatedAt = time.Now()
	err = s.transferRepository.SaveTransfer(transfer)
	if err != nil {
		return err
	}

	accountOrigin.Balance, accountDestination.Balance = accountOrigin.Balance - transfer.Amount, accountDestination.Balance + transfer.Amount
	s.accountService.UpdateAccount(accountOrigin)
	s.accountService.UpdateAccount(accountDestination)

	return nil
}