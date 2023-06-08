package core

import (
	"database/sql"
	"time"

	"github.com/daniial79/Banking-API/src/errs"
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

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	FindAllSql := "SELECT * FROM customers"

	rows, err := d.client.Query(FindAllSql)

	if err != nil {
		// fmt.Println(err)
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
			// fmt.Println(err)
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

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.DateofBirth, &c.Status, &c.Zipcode)

	if err != nil {
		if err == sql.ErrNoRows {
			// fmt.Println("ERROR: row not found")
			return nil, errs.NewNotFoundErr("customer not found")
		} else {
			// fmt.Println("ERROR: something went wrong during scan")
			return nil, errs.NewUnexpectedDbErr("Unexpected database error")
		}
	}

	return &c, nil

}
