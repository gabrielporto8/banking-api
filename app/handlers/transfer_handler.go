package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gabrielporto8/banking-api/app/models"
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.transferService.GetTransfers())
}

func (h TransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transfer models.Transfer

	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		log.Printf("Error decoding the body request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}

	err = h.transferService.CreateTransfer(&transfer)
	if err != nil {
		log.Printf("Error creating the transfer: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating the transfer."))
		return
	}

	json.NewEncoder(w).Encode(transfer)
}