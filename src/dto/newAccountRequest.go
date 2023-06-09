package dto

import (
	"strings"

	"github.com/daniial79/Banking-API/src/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000.00 {
		return errs.NewValidationErr("To open new account you need to deposit atleas t 5000.00")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationErr("Account type should be saving or checking")
	}

	return nil
}
