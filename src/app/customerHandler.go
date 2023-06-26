package app

import (
	"encoding/json"
	"net/http"

	"github.com/daniial79/Banking-API/src/dto"
	"github.com/daniial79/Banking-API/src/service"
	"github.com/gorilla/mux"
)

// Customer Primary Adapter
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) CreateNewCustomer(w http.ResponseWriter, r *http.Request) {
	var request dto.NewCustomerRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
		return
	}

	response, appErr := ch.service.NewCustomer(request)
	if err != nil {
		writeResponse(w, appErr.StatusCode, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusCreated, response)
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		writeResponse(w, err.StatusCode, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["customer_id"]
	customer, err := ch.service.GetCustomerById(id)

	if err != nil {
		writeResponse(w, err.StatusCode, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customer)
}
