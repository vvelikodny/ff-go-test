package errors

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPError(t *testing.T) {
	rr := httptest.NewRecorder()
	HTTPError(rr, "err", http.StatusInternalServerError)

	var r *HTTPErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&r))
	require.Equal(t, http.StatusInternalServerError, r.Code)
	require.Equal(t, "err", r.Message)
}
