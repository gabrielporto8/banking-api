package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gabrielporto8/stone-challenge/app/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var (
	accounts map[int64]*models.Account = make(map[int64]*models.Account)
	accountsLastID int64

	transfers map[int64]*models.Transfer = make(map[int64]*models.Transfer)
	transfersLastID int64
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(accounts)
		case http.MethodPost:	
			var account models.Account

			err := json.NewDecoder(r.Body).Decode(&account)
			if err != nil {
				log.Printf("Error decoding the body request: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("invalid request"))
				return
			}

			hashed, _ := bcrypt.GenerateFromPassword([]byte(account.Secret), 8)
	
			account.ID = accountsLastID
			account.Secret = string(hashed)
			account.CreatedAt = time.Now()
			accounts[accountsLastID] = &account
			accountsLastID++
	
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(account)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	r.HandleFunc("/accounts/{id}/balance", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		uriVars := mux.Vars(r)
		accountId := uriVars["id"]
		
		ID, err := strconv.ParseUint(accountId, 10, 64)
		if err != nil {
			errMsg := fmt.Sprintf("Error when parsing the given ID.")
			log.Println(errMsg)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}

		account, ok := accounts[int64(ID)]
		if !ok {
			errMsg := fmt.Sprintf("Account not found for the ID %v", ID)
			log.Println(errMsg)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errMsg))
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(account.Balance)
	})

	r.HandleFunc("/transfers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(transfers)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		var transfer models.Transfer

		err := json.NewDecoder(r.Body).Decode(&transfer)
		if err != nil {
			log.Printf("Error decoding the body request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request"))
			return
		}

		if transfer.AccountOriginID == transfer.AccountDestinationID {
			errMsg := fmt.Sprintf("Origin and Destination accounts must have different IDs")
			log.Println(errMsg)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}

		if transfer.Amount <= 0 {
			errMsg := fmt.Sprintf("The %v amount is not valid. The amount must be positive.", transfer.Amount)
			log.Println(errMsg)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}

		accountOrigin, ok := accounts[transfer.AccountOriginID]
		if !ok {
			errMsg := fmt.Sprintf("Origin account not found for the ID %v", transfer.AccountOriginID)
			log.Println(errMsg)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errMsg))
			return
		}

		accountDestination, ok := accounts[transfer.AccountDestinationID]
		if !ok {
			errMsg := fmt.Sprintf("Origin account not found for the ID %v", transfer.AccountDestinationID)
			log.Println(errMsg)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errMsg))
			return
		}

		if accountOrigin.Balance < transfer.Amount {
			errMsg := fmt.Sprintf("Origin account does not have sufficient amount")
			log.Println(errMsg)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errMsg))
			return
		}

		transfer.ID = transfersLastID
		transfers[transfersLastID] = &transfer
		transfersLastID++

		accountOrigin.Balance, accountDestination.Balance = accountOrigin.Balance - transfer.Amount, accountDestination.Balance + transfer.Amount

		accounts[accountOrigin.ID] = accountOrigin
		accounts[accountDestination.ID] = accountDestination

		json.NewEncoder(w).Encode(transfer)
	}).Methods("GET", "POST")
	
	log.Fatal(http.ListenAndServe(":8080", r))
}