package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"

	"github.com/vvelikodny/ff-go-test/api/errors"
	"github.com/vvelikodny/ff-go-test/api/services"
)

var (
	// Possible values for activityType field of deviceCheckRequest
	activityTypeList = []string{"SIGNUP", "LOGIN", "PAYMENT", "CONFIRMATION"}
)

// App represents main app with context
type App struct {
	sessionService services.SessionService
	Router         *mux.Router
}

// SessionService returns sessionService service
func (app *App) SessionService() services.SessionService {
	return app.sessionService
}

// NewApp creates new app with dependencies
func NewApp(service services.SessionService) *App {
	app := &App{sessionService: service}
	app.init()
	return app
}

func (app *App) init() {
	// register custom validator
	govalidator.CustomTypeTagMap.Set("activityType", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		value := i.(string)
		return govalidator.IsIn(value, activityTypeList...) || strings.HasPrefix(value, "_")
	}))

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/isgood", app.isGoodHandler).Methods(http.MethodPost)
}

// Run starts application and binds to port
func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

// deviceCheckRequest represents list of items of check request to universal SDK
type deviceCheckRequestList []deviceCheckRequest

// deviceCheckRequest represents item of check request to universal SDK
type deviceCheckRequest struct {
	CheckType       string `json:"checkType" valid:"in(DEVICE|BIOMETRIC|COMBO)~checkType should be either DEVICE|BIOMETRIC|COMBO,required"`
	ActivityType    string `json:"activityType" valid:"activityType~activityType should start with '_' or equal either SIGNUP|LOGIN|PAYMENT|CONFIRMATION,required"`
	CheckSessionKey string `json:"checkSessionKey" valid:"required"`
}

type deviceCheckResponse struct {
	Puppy bool `json:"puppy"`
}

// isGoodHandler
func (app *App) isGoodHandler(w http.ResponseWriter, r *http.Request) {
	var n deviceCheckRequestList
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		errors.HTTPError(w, fmt.Sprintf("JSON parsing error: %v", err), http.StatusBadRequest)
		return
	}

	for _, device := range n {
		if _, err := govalidator.ValidateStruct(device); err != nil {
			errors.HTTPError(w, fmt.Sprintf("New entity validation error: %v", err), http.StatusBadRequest)
			return
		}

		if !app.sessionService.Register(device.CheckSessionKey) {
			errors.HTTPError(w, fmt.Sprintf("Session key already registered: %v", device.CheckSessionKey), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&deviceCheckResponse{Puppy: true}); err != nil {
		errors.HTTPError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
