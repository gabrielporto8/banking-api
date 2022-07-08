package responses

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	WriteResponse(w, code, map[string]string {
		"Error": message,
	})
}