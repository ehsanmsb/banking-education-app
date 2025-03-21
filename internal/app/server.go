package app

import (
	"github.com/ehsanmsb/banking-education-app/internal/domain"
	"github.com/ehsanmsb/banking-education-app/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var ServerPort = ":9090"

type CustomerHandlers struct {
	service service.CustomerService
}

func StartServer() {
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// define route for get all users
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	// starting server on port 8000
	err := http.ListenAndServe(ServerPort, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
