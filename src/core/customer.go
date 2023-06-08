package core

import "github.com/daniial79/Banking-API/src/errs"

// Customer Core Object
type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

// Customer Secondary Port
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
}
