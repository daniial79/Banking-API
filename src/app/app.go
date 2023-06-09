package app

import (
	"net/http"

	"github.com/daniial79/Banking-API/src/config"
	"github.com/daniial79/Banking-API/src/core"
	"github.com/daniial79/Banking-API/src/logger"
	"github.com/daniial79/Banking-API/src/service"
	"github.com/gorilla/mux"
)

func Start() {

	//loading configurations
	config.InitAppConfig()

	router := mux.NewRouter()

	//wiring-up the application core
	customerHandler := CustomerHandler{service.NewDefaultCustomerService(core.NewCustomerRepositoryDb())}

	//routings
	router.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandler.GetCustomerById).Methods(http.MethodGet)

	if err := http.ListenAndServe(config.GetServerAddr(), router); err != nil {
		logger.Error(err.Error())
	}

}
