package service

import (
	"github.com/daniial79/Banking-API/src/core"
	"github.com/daniial79/Banking-API/src/dto"
	"github.com/daniial79/Banking-API/src/errs"
)

// Customer Primary Port
type CustomerService interface {
	NewCustomer(dto.NewCustomerRequest) (*dto.NewCustomerResponse, *errs.AppError)
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

func (s DefaultCustomerService) NewCustomer(req dto.NewCustomerRequest) (*dto.NewCustomerResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	customerCoreObject := core.Customer{
		Id:          "",
		Name:        req.Name,
		DateofBirth: req.DateOfbirth,
		City:        req.City,
		Zipcode:     req.Zipcode,
		Status:      "1",
	}

	newCustomer, err := s.repo.Save(customerCoreObject)
	if err != nil {
		return nil, err
	}

	response := newCustomer.ToNewCustomerResponseDto()

	return &response, nil
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
		response = append(response, customer.ToCustomerResponseDto())
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToCustomerResponseDto()

	return &response, nil

}
