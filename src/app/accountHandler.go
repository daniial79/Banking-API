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
		customerId = mux.Vars(r)["customer_id"]
	)
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
		return
	}

	request.CustomerId = customerId
	response, appError := ah.service.NewAccount(request)

	if appError != nil {
		writeResponse(w, appError.StatusCode, appError.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, response)
}

func (ah *AccountHandler) FetchMyAccounts(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)["customer_id"]

	response, err := ah.service.FetchAllAccounts(customerId)

	if err != nil {
		writeResponse(w, err.StatusCode, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, response)
}

func (ah *AccountHandler) FetchAccountById(w http.ResponseWriter, r *http.Request) {
	accountId := mux.Vars(r)["account_id"]
	response, err := ah.service.FetchAccountById(accountId)

	if err != nil {
		writeResponse(w, err.StatusCode, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, response)
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.NewTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		request.AccountId = accountId
		request.CustomerId = customerId

		account, appError := h.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.StatusCode, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}

}

func (ah AccountHandler) FetchAllAccountTransactions(w http.ResponseWriter, r *http.Request) {
	accountId := mux.Vars(r)["account_id"]
	transactionType := r.URL.Query().Get("type")
	response, err := ah.service.FetchAllAccountTransactions(accountId, transactionType)

	if err != nil {
		writeResponse(w, err.StatusCode, err.AsMessage())
		return
	}

	writeResponse(w, err.StatusCode, response)
}
