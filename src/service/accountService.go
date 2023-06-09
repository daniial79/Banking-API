package service

import (
	"time"

	"github.com/daniial79/Banking-API/src/core"
	"github.com/daniial79/Banking-API/src/dto"
	"github.com/daniial79/Banking-API/src/errs"
)

// Account Primary Port
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

// Account Default Service Primary Adapter
type DefaultAccountService struct {
	repo core.AccountRepository
}

func NewAccountService(repo core.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	coreAccountObject := core.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	NewAccount, err := s.repo.Save(coreAccountObject)
	if err != nil {
		return nil, err
	}

	newAccountResponse := NewAccount.ToNewAccountResponseDto()

	return &newAccountResponse, nil
}
