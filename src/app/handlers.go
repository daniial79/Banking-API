package app

import (
	"encoding/json"
	"net/http"

	"github.com/daniial79/Banking-API/src/service"
)

// Customer Primary Adapter
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
