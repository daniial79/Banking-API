package app

import (
	"encoding/json"
	"net/http"

	"github.com/daniial79/Banking-API/src/dto"
	"github.com/daniial79/Banking-API/src/service"
	"github.com/gorilla/mux"
)

// Account Primary Adapter
type AccountHandler struct {
	service service.AccountService
}

func (ah *AccountHandler) CreateNewAccount(w http.ResponseWriter, r *http.Request) {
	var (
		request    dto.NewAccountRequest
		customerId = mux.Vars(r)["id"]
	)
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	request.CustomerId = customerId
	response, appError := ah.service.NewAccount(request)

	if appError != nil {
		WriteResponse(w, appError.StatusCode, appError.AsMessage())
		return
	}

	WriteResponse(w, http.StatusCreated, response)
}
