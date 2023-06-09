package service

import (
	"github.com/daniial79/Banking-API/src/core"
	"github.com/daniial79/Banking-API/src/dto"
	"github.com/daniial79/Banking-API/src/errs"
)

// Customer Primary Port
type CustomerService interface {
	//TODO implement create customer + dto req and response
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError)
}

// Customer Service Primary Adapter
type DefaultCustomerService struct {
	repo core.CustomerRepository
}

func NewDefaultCustomerService(repository core.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)

	for _, customer := range customers {
		response = append(response, customer.ToDto())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil

}
