package app

import (
	"log"
	"net/http"

	"github.com/gabrielporto8/banking-api/pkg/handlers"
	"github.com/gabrielporto8/banking-api/pkg/middlewares"
	"github.com/gabrielporto8/banking-api/pkg/repositories"
	"github.com/gabrielporto8/banking-api/pkg/services"
	"github.com/gorilla/mux"
)

var (
	accountRepository *repositories.AccountRepository
	transferRepository *repositories.TransferRepository

	accountService *services.AccountService
	transferService *services.TransferService
	jwtService *services.JWTService
	authService *services.AuthService

	accountHandler *handlers.AccountHandler
	transferHandler *handlers.TransferHandler
	authHandler *handlers.AuthHandler
)

func Run() {
	setHandlers()
	router := setRoutes()
	
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setHandlers() {
	accountRepository = repositories.NewAccountRepository()
	accountService = services.NewAccountService(accountRepository)
	accountHandler = handlers.NewAccountHandler(accountService)

	transferRepository = repositories.NewTransferRepository()
	transferService = services.NewTransferService(transferRepository, accountService)
	transferHandler = handlers.NewTransferHandler(transferService)

	jwtService = services.NewJWTService()
	authService = services.NewAuthService(accountService, jwtService)
	authHandler = handlers.NewAuthHandler(authService)
}

func setRoutes() *mux.Router {
	r := mux.NewRouter()
	
	secure := r.PathPrefix("/transfers").Subrouter()
	secure.Use(middlewares.AuthMiddleware)
	secure.HandleFunc("", transferHandler.GetTransfers).Methods("GET")
	secure.HandleFunc("", transferHandler.CreateTransfer).Methods("POST")

	r.HandleFunc("/accounts", accountHandler.GetAccounts).Methods("GET")
	r.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalance).Methods("GET")
	
	r.HandleFunc("/login", authHandler.GenerateToken).Methods("POST")

	return r
}