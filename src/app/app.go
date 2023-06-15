package app

import (
	"net/http"
	"time"

	"github.com/daniial79/Banking-API/src/config"
	"github.com/daniial79/Banking-API/src/core"
	"github.com/daniial79/Banking-API/src/logger"
	"github.com/daniial79/Banking-API/src/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	//loading initial configurations
	config.InitAppConfig()
	dbClient := getNewPullConnectionDb()
	router := mux.NewRouter()

	//wiring-up the application core
	customerRepositoryDb := core.NewCustomerRepositoryDb(dbClient)
	ch := CustomerHandler{service.NewDefaultCustomerService(customerRepositoryDb)}

	accountRepositoryDb := core.NewAccountRepositoryDb(dbClient)
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	//routings
	router.
		HandleFunc("/customers", ch.CreateNewCustomer).
		Methods(http.MethodPost).
		Name("CreateCustomer")

	router.
		HandleFunc("/customers", ch.GetAllCustomers).
		Methods(http.MethodGet).
		Name("GetCustomers")

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomerById).
		Methods(http.MethodGet).
		Name("GetByCustomerId")

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/accounts", ah.CreateNewAccount).
		Methods(http.MethodPost).
		Name("CreateAccount")

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/accounts/{account_id:[0-9]+}", ah.FetchAccountById).
		Methods(http.MethodGet).
		Name("GetByAccountId")

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/accounts", ah.FetchMyAccounts).
		Methods(http.MethodGet).
		Name("GetAllCustomerAccounts")

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]}/transactions", ah.FetchAllAccountTransactions).
		Methods(http.MethodGet).
		Name("GetAllAccountTransactions")

	if err := http.ListenAndServe(config.GetServerAddr(), router); err != nil {
		logger.Error(err.Error())
	}

}

func getNewPullConnectionDb() *sqlx.DB {
	db, err := sqlx.Open(
		config.GetDbDialect(),
		config.GetDatabaseSourceName(),
	)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
