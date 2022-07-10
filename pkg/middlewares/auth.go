package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gabrielporto8/banking-api/pkg/responses"
	"github.com/gabrielporto8/banking-api/pkg/services"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)
		
		jwtService := services.NewJWTService()
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			log.Printf("Error when validating the token: %v", err.Error())
			responses.WriteErrorResponse(w, err.Code, err.Error())
			return
		}

		r.Header.Add("Authenticated-CPF", claims.Cpf)
		next.ServeHTTP(w, r)
	})
}

func getTokenFromRequest(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")

	if len(strings.Split(tokenString, " ")) == 2 {
		return strings.Split(tokenString, " ")[1]
	}

	return ""
}