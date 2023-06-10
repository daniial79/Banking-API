package core

import (
	"database/sql"
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

func (d AccountRepositoryDb) FindAllCustomerAccounts(customerId string) ([]Account, *errs.AppError) {
	accounts := make([]Account, 0)
	FindAllSql := "SELECT * FROM accounts WHERE customer_id = ?"

	err := d.client.Select(&accounts, FindAllSql, customerId)
	if err != nil {
		logger.Error("Error while retrieving all the accounts associated to specific customer: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected database error")
	}

	if len(accounts) == 0 {
		return nil, errs.NewNotFoundErr("There is no account registered with this customer id")
	}

	return accounts, nil
}

func (d AccountRepositoryDb) FindById(accountId string) (*Account, *errs.AppError) {
	var account Account
	findByIdSql := "SELECT * FROM accounts WHERE account_id = ?"

	err := d.client.Get(&account, findByIdSql, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundErr("There is no account registered with this account id")
		}
		logger.Error("Error while retrieving account: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected database error")
	}

	return &account, nil

}
