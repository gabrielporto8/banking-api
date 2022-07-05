package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gabrielporto8/stone-challenge/app/models"
	"github.com/gabrielporto8/stone-challenge/app/services"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	accountService *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := h.accountService.GetAccounts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func (h AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Printf("Error decoding the body request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}

	created := h.accountService.CreateAccount(&account)
	if !created {
		log.Printf("Error creating the account.")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating the account."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (h AccountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	uriVars := mux.Vars(r)
	accountId := uriVars["id"]
	
	ID, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Error when parsing the given ID.")
		log.Println(errMsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errMsg))
		return
	}

	balance, ok := h.accountService.GetBalance(ID)
	if !ok {
		errMsg := fmt.Sprintf("Account not found for the ID %v", ID)
		log.Println(errMsg)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errMsg))
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}