package handlers

import (
	"errors"
	"net/http"

	"github.com/gabrielporto8/banking-api/pkg/errs"
)

var ErrHeaderNotFound = errors.New("Failed when getting the data from authenticated user.")

func getAuthenticatedCpfFromRequestHeader(r *http.Request) (string, *errs.AppError) {
	cpf := r.Header.Get("Authenticated-CPF")
	if len(cpf) == 0 {
		return "", errs.NewUnauthorizedError(ErrHeaderNotFound)
	}

	return cpf, nil
}