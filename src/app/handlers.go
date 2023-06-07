package app

import (
	"github.com/daniial79/Banking-API/src/service"
)

// Customer Primary Adapter
type CustomerHandler struct {
	service service.CustomerService
}
