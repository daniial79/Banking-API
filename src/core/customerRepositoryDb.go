package core

import (
	"database/sql"
	"time"

	"github.com/daniial79/Banking-API/src/errs"
	"github.com/daniial79/Banking-API/src/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Customer Repository Db Secondary Adapter
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	db, err := sqlx.Open("mysql", "root:13454779d@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDb{db}
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
