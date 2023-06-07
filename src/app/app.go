package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type responseSample struct {
	Message string `json:"message"`
}

func greetController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(responseSample{
		Message: "Hello from gorilla multiplexer!",
	})
}

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", greetController)

	fmt.Println("server is up and running on port 8000...")
	log.Fatalln(http.ListenAndServe(":8000", router))

}
