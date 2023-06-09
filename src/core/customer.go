package core

import (
	"github.com/daniial79/Banking-API/src/dto"
	"github.com/daniial79/Banking-API/src/errs"
)

// Customer Core Object
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) ToNewCustomerResponseDto() dto.NewCustomerResponse {
	return dto.NewCustomerResponse{
		CustomerId: c.Id,
	}
}

// Customer Secondary Port
type CustomerRepository interface {
	Save(Customer) (*Customer, *errs.AppError)
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}

func (c Customer) setStatusAsText() string {
	stat := "active"
	if c.Status == "0" {
		stat = "inactive"
	}
	return stat
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.setStatusAsText(),
	}
}
