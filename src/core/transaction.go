package core

import "github.com/daniial79/Banking-API/src/dto"

const WITHDRAWAL = "withdrawal"

// Transaction Core Object
type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) ToTransactionResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		Amount:          t.Amount,
		TransactionType: t.AccountId,
		TransactionDate: t.TransactionDate,
	}
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
