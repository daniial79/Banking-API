package app

import (
	"net/http"

	"github.com/daniial79/Banking-API/src/service"
	"github.com/gorilla/mux"
)

// Customer Primary Adapter
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		WriteResponse(w, err.StatusCode, err.AsMessage())
		return
	}

	WriteResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	customer, err := ch.service.GetCustomerById(id)

	if err != nil {
		WriteResponse(w, err.StatusCode, err.AsMessage())
		return
	}

	WriteResponse(w, http.StatusOK, customer)
}
