package errors

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPError(t *testing.T) {
	rr := httptest.NewRecorder()
	HTTPError(rr, "err", http.StatusInternalServerError)

	var r *HTTPErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&r))
	require.Equal(t, http.StatusInternalServerError, r.Code)
	require.Equal(t, "err", r.Message)
}
