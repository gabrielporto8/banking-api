package services

import (
	"time"

	"github.com/gabrielporto8/stone-challenge/app/models"
	"github.com/gabrielporto8/stone-challenge/app/repositories"
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

func (s TransferService) CreateTransfer(transfer *models.Transfer) bool {
	if transfer.AccountOriginID == transfer.AccountDestinationID {
		return false
	}

	if transfer.Amount <= 0 {
		return false
	}

	accountOrigin, ok := s.accountService.GetAccountByID(transfer.AccountOriginID)
	if !ok {
		return false
	}

	accountDestination, ok := s.accountService.GetAccountByID(transfer.AccountDestinationID)
	if !ok {
		return false
	}

	if accountOrigin.Balance < transfer.Amount {
		return false
	}

	transfer.CreatedAt = time.Now()
	success := s.transferRepository.SaveTransfer(transfer)
	if !success {
		return false
	}

	accountOrigin.Balance, accountDestination.Balance = accountOrigin.Balance - transfer.Amount, accountDestination.Balance + transfer.Amount
	s.accountService.UpdateAccount(accountOrigin)
	s.accountService.UpdateAccount(accountDestination)
	
	return true
}