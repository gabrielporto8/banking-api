package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gabrielporto8/banking-api/app/models"
	"github.com/gabrielporto8/banking-api/app/responses"
	"github.com/gabrielporto8/banking-api/app/services"
)

type TransferHandler struct {
	transferService *services.TransferService
}

func NewTransferHandler(transferService *services.TransferService) *TransferHandler {
	return &TransferHandler{
		transferService: transferService,
	}
}

func (h TransferHandler) GetTransfers(w http.ResponseWriter, r *http.Request) {
	cpf, err := getAuthenticatedCpfFromRequestHeader(r)
	if err != nil {
		log.Printf("Error when getting CPF from header: %v", err.Error())
		responses.WriteErrorResponse(w, err.Code, err.Error())
		return
	} 

	transfers, err := h.transferService.GetTransfersByCPF(cpf)
	if err != nil {
		log.Printf("Error when getting the transfers: %v", err.Error())
		responses.WriteErrorResponse(w, err.Code, err.Error())
		return
	}
	
	responses.WriteResponse(w, http.StatusAccepted, transfers)
}

func (h TransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	cpf, appError := getAuthenticatedCpfFromRequestHeader(r)
	if appError != nil {
		log.Printf("Error when getting CPF from header: %v", appError.Error())
		responses.WriteErrorResponse(w, appError.Code, appError.Error())
		return
	}
	
	var transfer models.Transfer
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		log.Printf("Error decoding the body request: %v", err)
		responses.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	appError = h.transferService.CreateTransferFromRequest(&transfer, cpf)
	if appError != nil {
		log.Printf("Error creating the transfer: %v", appError.Error())
		responses.WriteErrorResponse(w, appError.Code, appError.Error())
		return
	}

	responses.WriteResponse(w, http.StatusAccepted, transfer)
}