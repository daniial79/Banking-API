package service

import "github.com/daniial79/Banking-API/src/core"

// Customer Primary Port
type CustomerService interface {
	GetAllCustomers() ([]core.Customer, error)
}

// Customer Service Primary Adapter
type DefaultCustomerService struct {
	repo core.CustomerRepository
}

func NewDefaultCustomerService(repository core.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetAllCustomers() ([]core.Customer, error) {
	return s.repo.FindAll()
}
