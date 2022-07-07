package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gabrielporto8/banking-api/app/models"
	"github.com/gabrielporto8/banking-api/app/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h AuthHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var authentication models.Authentication

	err := json.NewDecoder(r.Body).Decode(&authentication)
	if err != nil {
		log.Printf("Error decoding the body request: %v", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	token, appError := h.authService.GenerateToken(&authentication)
	if appError != nil {
		log.Printf("Authentication error: %v", appError.Error())
		writeResponse(w, appError.Code, appError.Error())
		return
	}

	writeResponse(w, http.StatusOK, token)
}
