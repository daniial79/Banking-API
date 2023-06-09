package core

import (
	"strconv"

	"github.com/daniial79/Banking-API/src/errs"
	"github.com/daniial79/Banking-API/src/logger"
	"github.com/jmoiron/sqlx"
)

// Account Repository Db Second Adapter
type AccountRepositoryDb struct {
	client *sqlx.DB
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	insertNewAccountSql := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(insertNewAccountSql, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while fetching last inserted account's id: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected database error")
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}
