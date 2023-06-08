package app

import (
	"log"
	"net/http"

	"github.com/daniial79/Banking-API/src/core"
	"github.com/daniial79/Banking-API/src/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring-up the application components
	customerHandler := CustomerHandler{service.NewDefaultCustomerService(core.NewCustomerRepositoryDb())}

	//routings
	router.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", customerHandler.GetCustomerById).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe(":8000", router))

}
