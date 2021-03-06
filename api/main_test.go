package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vvelikodny/ff-go-test/api/errors"
	"github.com/vvelikodny/ff-go-test/api/services"
)

func TestAppGetSessionService(t *testing.T) {
	ss := services.NewInMemSessionService()
	app := NewApp(ss)

	require.Equal(t, ss, app.SessionService())
}

func TestBadCheckRequestError(t *testing.T) {
	app := NewApp(services.NewInMemSessionService())

	req, _ := http.NewRequest(http.MethodPost, "/isgood", bytes.NewBuffer([]byte(``)))
	response := executeRequest(t, app, req)
	require.Equal(t, http.StatusInternalServerError, response.Code)

	var m errors.HTTPErrorResponse
	require.NoError(t, json.NewDecoder(response.Body).Decode(&m))

	require.Equal(t, http.StatusBadRequest, m.Code)
	require.Contains(t, m.Message, "JSON parsing error")
}

func TestEmptyCheckRequestError(t *testing.T) {
	app := NewApp(services.NewInMemSessionService())

	req, _ := http.NewRequest(http.MethodPost, "/isgood", bytes.NewBuffer([]byte(`[{}]`)))
	response := executeRequest(t, app, req)
	require.Equal(t, http.StatusInternalServerError, response.Code)

	var m errors.HTTPErrorResponse
	require.NoError(t, json.NewDecoder(response.Body).Decode(&m))

	require.Equal(t, m.Code, http.StatusBadRequest)
	require.Contains(t, m.Message, "checkType: non zero value required")
	require.Contains(t, m.Message, "activityType: non zero value require")
	require.Contains(t, m.Message, "checkSessionKey: non zero value required")
}

func TestWrongCheckType(t *testing.T) {
	app := NewApp(services.NewInMemSessionService())

	req, _ := http.NewRequest(http.MethodPost, "/isgood", bytes.NewBuffer([]byte(`[{"checkType": "DEVICE1","activityType": "LOGIN","checkSessionKey": "123"}]`)))
	response := executeRequest(t, app, req)
	require.Equal(t, http.StatusInternalServerError, response.Code)

	var m errors.HTTPErrorResponse
	require.NoError(t, json.NewDecoder(response.Body).Decode(&m))
	require.Contains(t, m.Message, "checkType should be either DEVICE|BIOMETRIC|COMBO")
}

func TestWrongActivityType(t *testing.T) {
	app := NewApp(services.NewInMemSessionService())

	req, _ := http.NewRequest(http.MethodPost, "/isgood", bytes.NewBuffer([]byte(`[{"checkType": "DEVICE","activityType": "LOGIN1","checkSessionKey": "123"}]`)))
	response := executeRequest(t, app, req)
	require.Equal(t, http.StatusInternalServerError, response.Code)

	var m errors.HTTPErrorResponse
	require.NoError(t, json.NewDecoder(response.Body).Decode(&m))
	require.Contains(t, m.Message, "activityType should start with '_' or equal either SIGNUP|LOGIN|PAYMENT|CONFIRMATION")
}

func TestActivityTypeUnderscore(t *testing.T) {
	app := NewApp(services.NewInMemSessionService())

	req, _ := http.NewRequest(http.MethodPost, "/isgood", bytes.NewBuffer([]byte(`[{"checkType": "DEVICE","activityType": "_LOGIN_1","checkSessionKey": "123"}]`)))
	response := executeRequest(t, app, req)
	require.Equal(t, http.StatusOK, response.Code)
}

func TestDuplicateSessionKey(t *testing.T) {
	app := NewApp(services.NewInMemSessionService())

	req, _ := http.NewRequest(http.MethodPost, "/isgood", bytes.NewBuffer([]byte(`[{"checkType": "DEVICE","activityType": "LOGIN","checkSessionKey": "123"}]`)))
	response := executeRequest(t, app, req)
	require.Equal(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest(http.MethodPost, "/isgood", bytes.NewBuffer([]byte(`[{"checkType": "DEVICE","activityType": "LOGIN","checkSessionKey": "123"}]`)))
	response = executeRequest(t, app, req)
	require.Equal(t, http.StatusInternalServerError, response.Code)

	var m errors.HTTPErrorResponse
	require.NoError(t, json.NewDecoder(response.Body).Decode(&m))
	require.Contains(t, m.Message, "Session key already registered: 123")

}

func executeRequest(t *testing.T, app *App, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}
