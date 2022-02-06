package app

import (
	"log"
	"net/http"

	"github.com/golang/Hexagonal-golangBancking/domain"
	"github.com/golang/Hexagonal-golangBancking/service"
	"github.com/gorilla/mux"
	//"google.golang.org/genproto/googleapis/rpc/status"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}