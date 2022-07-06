package main

import (
	"log"
	"net/http"

	"github.com/gabrielporto8/banking-api/app/handlers"
	"github.com/gabrielporto8/banking-api/app/middlewares"
	"github.com/gabrielporto8/banking-api/app/repositories"
	"github.com/gabrielporto8/banking-api/app/services"
	"github.com/gorilla/mux"
)

func main() {
	accountRepository := repositories.NewAccountRepository()
	accountService := services.NewAccountService(accountRepository)
	accountHandler := handlers.NewAccountHandler(accountService)

	transferRepository := repositories.NewTransferRepository()
	transferService := services.NewTransferService(transferRepository, accountService)
	transferHandler := handlers.NewTransferHandler(transferService)

	authRepository := repositories.NewAuthRepository()

	jwtService := services.NewJWTService()
	authService := services.NewAuthService(authRepository, accountRepository, jwtService)
	authHandler := handlers.NewAuthHandler(authService)

	r := mux.NewRouter()
	
	secure := r.PathPrefix("/transfers").Subrouter()
	secure.Use(middlewares.AuthMiddleware)
	secure.HandleFunc("", transferHandler.GetTransfers).Methods("GET")
	secure.HandleFunc("", transferHandler.CreateTransfer).Methods("POST")

	r.HandleFunc("/accounts", accountHandler.GetAccounts).Methods("GET")
	r.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalance).Methods("GET")
	
	r.HandleFunc("/login", authHandler.GenerateToken).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", r))
}