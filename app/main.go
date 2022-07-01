package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gabrielporto8/stone-challenge/app/models"
)

var (
	accounts map[int64]*models.Account = make(map[int64]*models.Account)
	accountsLastID int64

	transfers map[int64]*models.Transfer = make(map[int64]*models.Transfer)
	transfersLastID int64
)

func main() {
	http.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
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
	
			account.ID = accountsLastID
			account.CreatedAt = time.Now()
			accounts[accountsLastID] = &account
			accountsLastID++
	
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(account)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}