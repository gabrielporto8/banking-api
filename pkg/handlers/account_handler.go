package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gabrielporto8/banking-api/pkg/models"
	"github.com/gabrielporto8/banking-api/pkg/responses"
	"github.com/gabrielporto8/banking-api/pkg/services"
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
	responses.WriteResponse(w, http.StatusAccepted, accounts)
}

func (h AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Printf("Error decoding the body request: %v", err)
		responses.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	appError := h.accountService.CreateAccount(&account)
	if appError != nil {
		log.Printf("Error creating the account: %v", appError.Error())
		responses.WriteErrorResponse(w, appError.Code, appError.Error())
		return
	}

	responses.WriteResponse(w, http.StatusAccepted, account)
}

func (h AccountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	uriVars := mux.Vars(r)
	accountId := uriVars["id"]

	ID, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Error when parsing the given ID.")
		log.Println(errMsg)
		responses.WriteErrorResponse(w, http.StatusBadRequest, errMsg)
		return
	}

	balance, appError := h.accountService.GetBalance(ID)
	if appError != nil {
		log.Printf("Error on getting the account balance: %v", appError.Error())
		responses.WriteErrorResponse(w, appError.Code, appError.Error())
		return
	}

	responses.WriteResponse(w, http.StatusAccepted, map[string]float64 {"Balance": balance})
}
