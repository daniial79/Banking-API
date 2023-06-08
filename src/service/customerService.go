package service

import (
	"github.com/daniial79/Banking-API/src/core"
	"github.com/daniial79/Banking-API/src/errs"
)

// Customer Primary Port
type CustomerService interface {
	GetAllCustomers() ([]core.Customer, *errs.AppError)
	GetCustomerById(id string) (*core.Customer, *errs.AppError)
}

// Customer Service Primary Adapter
type DefaultCustomerService struct {
	repo core.CustomerRepository
}

func NewDefaultCustomerService(repository core.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (s DefaultCustomerService) GetAllCustomers() ([]core.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerById(id string) (*core.Customer, *errs.AppError) {
	return s.repo.FindById(id)
}
