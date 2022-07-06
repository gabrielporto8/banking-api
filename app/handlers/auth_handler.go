package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gabrielporto8/stone-challenge/app/models"
	"github.com/gabrielporto8/stone-challenge/app/services"
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}
	
	token, err := h.authService.GenerateToken(&authentication)
	if err != nil {
		log.Printf("Authentication error: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
