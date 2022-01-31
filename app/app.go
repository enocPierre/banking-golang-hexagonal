package app

import (
	"log"
	"net/http"

	"github.com/golang/Hexagonal-golangBancking/domain"
	"github.com/golang/Hexagonal-golangBancking/service"
	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}


	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}