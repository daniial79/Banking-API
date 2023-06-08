package core

import (
	"database/sql"
	"time"

	"github.com/daniial79/Banking-API/src/errs"
	"github.com/daniial79/Banking-API/src/logger"
	_ "github.com/go-sql-driver/mysql"
)

// Customer Repository Db Secondary Adapter
type CustomerRepositoryDb struct {
	client *sql.DB
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	db, err := sql.Open("mysql", "root:13454779d@tcp(localhost:3306)/banking")
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

	rows, err := d.client.Query(FindAllSql)

	if err != nil {
		logger.Error("Error while querying for customers: " + err.Error())
		return nil, errs.NewUnexpectedDbErr("Unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer

		err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.City,
			&c.DateofBirth,
			&c.Zipcode,
			&c.Status,
		)

		if err != nil {
			logger.Error("Error while scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedDbErr("Unexpected database error")
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	FindByIdQuery := "SELECT * FROM customers WHERE customer_id = ?"

	row := d.client.QueryRow(FindByIdQuery, id)
	var c Customer

	err := row.Scan(
		&c.Id,
		&c.Name,
		&c.City,
		&c.DateofBirth,
		&c.Status,
		&c.Zipcode,
	)

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
