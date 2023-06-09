package core

import (
	"database/sql"

	"github.com/daniial79/Banking-API/src/errs"
	"github.com/daniial79/Banking-API/src/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Customer Repository Db Secondary Adapter
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	FindAllSql := "SELECT * FROM customers"

	if status == "1" {
		FindAllSql += " WHERE status = 1"
	} else if status == "0" {
		FindAllSql += " WHERE status = 0"
	}

	customers := make([]Customer, 0)
	err := d.client.Select(&customers, FindAllSql)

	if err != nil {
		logger.Error("Error while querying for customers: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	FindByIdQuery := "SELECT * FROM customers WHERE customer_id = ?"

	var c Customer
	err := d.client.Get(&c, FindByIdQuery, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundErr("customer not found")
		} else {
			logger.Error("Error while scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedDbErr("Unexpected database error")
		}
	}

	return &c, nil

}
