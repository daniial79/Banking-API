package core

import (
	"github.com/daniial79/Banking-API/src/dto"
	"github.com/daniial79/Banking-API/src/errs"
)

// Account Core Object
type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

func (a Account) ToAccountDto() dto.AccountResponse {
	return dto.AccountResponse{
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.setStatusAsText(),
	}
}

func (a Account) setStatusAsText() string {
	stat := "active"
	if a.Status == "0" {
		stat = "inactive"
	}
	return stat
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}

// Account Secondary Port
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindAllCustomerAccounts(string) ([]Account, *errs.AppError)
	FindById(customerId string) (*Account, *errs.AppError)
	SaveTransaction(t Transaction) (*Transaction, *errs.AppError)
	GetTransactions(accountId, transactionType string) ([]Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}
