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
	FetchAccountById(string) (*dto.AccountResponse, *errs.AppError)
	FetchAllAccounts(customerId string) ([]dto.AccountResponse, *errs.AppError)
	MakeTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
	FetchAllAccountTransactions(accountId, transactionType string) ([]dto.TransactionResponse, *errs.AppError)
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

func (s DefaultAccountService) FetchAllAccounts(customerId string) ([]dto.AccountResponse, *errs.AppError) {
	coreObjectAccounts, err := s.repo.FindAllCustomerAccounts(customerId)

	if err != nil {
		return nil, err
	}

	accounts := make([]dto.AccountResponse, 0)

	for _, account := range coreObjectAccounts {
		accounts = append(accounts, account.ToAccountDto())
	}

	return accounts, nil
}

func (s DefaultAccountService) FetchAccountById(accountId string) (*dto.AccountResponse, *errs.AppError) {
	accountCoreObject, err := s.repo.FindById(accountId)

	if err != nil {
		return nil, err
	}

	accountResponse := accountCoreObject.ToAccountDto()

	return &accountResponse, nil

}

func (s DefaultAccountService) MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationErr("Insufficient balance in the account")
		}
	}

	t := core.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToNewTransactionResponseDto()
	return &response, nil
}

func (s DefaultAccountService) FetchAllAccountTransactions(accountId, transactionType string) ([]dto.TransactionResponse, *errs.AppError) {
	coreTransactionObjects, err := s.repo.GetTransactions(accountId, transactionType)
	if err != nil {
		return nil, err
	}

	response := make([]dto.TransactionResponse, 0)

	for _, transaction := range coreTransactionObjects {
		response = append(response, transaction.ToTransactionResponseDto())
	}

	return response, nil
}
