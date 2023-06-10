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
	router.HandleFunc("/customers", ch.CreateNewCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.GetCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}/account", ah.CreateNewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id:[0-9]+}/accounts", ah.FetchMyAccountsId)

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
