package services

import (
	"errors"
	"time"

	"github.com/gabrielporto8/banking-api/pkg/errs"
	"github.com/gabrielporto8/banking-api/pkg/models"
	"github.com/gabrielporto8/banking-api/pkg/repositories"
)

var (
	ErrAccountOriginNotFound = errors.New("Account origin not found.")
	ErrAccountDestinationNotFound = errors.New("Account destination not found.")
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

func (s TransferService) GetTransfersByCPF(cpf string) ([]models.Transfer, *errs.AppError) {
	account, err := s.accountService.GetAccountByCPF(cpf)
	if err != nil {
		return nil, err
	}
	return s.GetTransfersByOriginID(account.ID), nil
}

func (s TransferService) GetTransfersByOriginID(ID int64) []models.Transfer {
	return s.transferRepository.GetTransfersByOriginID(ID)
}

func (s TransferService) CreateTransferFromRequest(transfer *models.Transfer, cpf string) *errs.AppError {
	err := s.setAccountOriginIDByCPF(transfer, cpf)
	if err != nil {
		return err
	}
	
	return s.CreateTransfer(transfer)
}

func (s TransferService) setAccountOriginIDByCPF(transfer *models.Transfer, cpf string) *errs.AppError {
	accountOrigin, err := s.accountService.GetAccountByCPF(cpf)
	if err != nil {
		return err
	}

	transfer.AccountOriginID = accountOrigin.ID

	return nil
}

func (s TransferService) CreateTransfer(transfer *models.Transfer) *errs.AppError {
	err := transfer.Validate()
	if err != nil {
		return err
	}

	_, err = s.accountService.GetAccountByID(transfer.AccountDestinationID)
	if err != nil {
		return errs.NewNotFoundError(ErrAccountDestinationNotFound)
	}

	accountOrigin, err := s.accountService.GetAccountByID(transfer.AccountOriginID)
	if err != nil {
		return errs.NewNotFoundError(ErrAccountOriginNotFound)
	}

	err = accountOrigin.ValidateTransferBalance(transfer.Amount)
	if err != nil {
		return err
	}

	transfer.CreatedAt = time.Now()
	err = s.transferRepository.SaveTransfer(transfer)
	if err != nil {
		return err
	}

	err = s.accountService.UpdateAccountBalanceFromTransfer(transfer)
	if err != nil {
		return err
	}

	return nil
}