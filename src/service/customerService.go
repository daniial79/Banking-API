package service

import (
	"github.com/daniial79/Banking-API/src/core"
	"github.com/daniial79/Banking-API/src/dto"
	"github.com/daniial79/Banking-API/src/errs"
)

// Customer Primary Port
type CustomerService interface {
	GetAllCustomers(status string) ([]core.Customer, *errs.AppError)
	GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError)
}

// Customer Service Primary Adapter
type DefaultCustomerService struct {
	repo core.CustomerRepository
}

func NewDefaultCustomerService(repository core.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]core.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil

}
