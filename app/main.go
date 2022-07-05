package main

import (
	"log"
	"net/http"

	"github.com/gabrielporto8/stone-challenge/app/handlers"
	"github.com/gabrielporto8/stone-challenge/app/repositories"
	"github.com/gabrielporto8/stone-challenge/app/services"
	"github.com/gorilla/mux"
)

func main() {
	accountRepository := repositories.NewAccountRepository()
	accountService := services.NewAccountService(accountRepository)
	accountHandler := handlers.NewAccountHandler(accountService)

	transferRepository := repositories.NewTransferRepository()
	transferService := services.NewTransferService(transferRepository, accountService)
	transferHandler := handlers.NewTransferHandler(transferService)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", accountHandler.GetAccounts).Methods("GET")
	r.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalance).Methods("GET")
	r.HandleFunc("/transfers", transferHandler.GetTransfers).Methods("GET")
	r.HandleFunc("/transfers", transferHandler.CreateTransfer).Methods("POST")

	r.HandleFunc("/login", handlers.GenerateToken)
	
	log.Fatal(http.ListenAndServe(":8080", r))
}