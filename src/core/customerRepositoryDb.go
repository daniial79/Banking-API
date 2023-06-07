package core

import (
	"database/sql"
	"fmt"
	"time"

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

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	FindAllSql := "SELECT * FROM customers"

	rows, err := d.client.Query(FindAllSql)

	if err != nil {
		fmt.Println(err)
		return nil, err
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
			fmt.Println(err)
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}
