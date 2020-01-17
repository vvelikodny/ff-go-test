package errors

import (
	"encoding/json"
	"net/http"
)

// HTTPErrorResponse represents common HTTP error
type HTTPErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HTTPError replies as response with specific message & HTTP code.
func HTTPError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(&HTTPErrorResponse{
		Code:    code,
		Message: err,
	})
}
