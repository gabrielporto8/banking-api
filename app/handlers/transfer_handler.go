package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gabrielporto8/banking-api/app/models"
	"github.com/gabrielporto8/banking-api/app/services"
)

var ErrHeaderNotFound = errors.New("failed when getting the data from authenticated user")

type TransferHandler struct {
	transferService *services.TransferService
}

func NewTransferHandler(transferService *services.TransferService) *TransferHandler {
	return &TransferHandler{
		transferService: transferService,
	}
}

func (h TransferHandler) GetTransfers(w http.ResponseWriter, r *http.Request) {
	cpf := w.Header().Get("cpf_authenticated")
	if len(cpf) == 0 {
		log.Printf("Error: %v", ErrHeaderNotFound)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid request"))
		return
	}

	transfers, err := h.transferService.GetTransfersByCPF(cpf)
	if err != nil {
		log.Printf("Error when getting the transfers: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid request"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transfers)
}

func (h TransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	cpf := w.Header().Get("cpf_authenticated")
	if len(cpf) == 0 {
		log.Printf("Error: %v", ErrHeaderNotFound)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid request"))
		return
	}
	
	var transfer models.Transfer
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		log.Printf("Error decoding the body request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}
	
	err = h.transferService.CreateTransfer(&transfer, cpf)
	if err != nil {
		log.Printf("Error creating the transfer: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating the transfer."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transfer)
}