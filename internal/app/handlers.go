package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()
	if r.Header.Get("content-type") == "application/xml" {
		w.Header().Set("content-type", "application/xml")
		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			log.Fatal("xml err: ", err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(customers)
		if err != nil {
			log.Println("getAllUsers err: ", err)
		}
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerId := params["customer_id"]
	customer, err := ch.service.GetCustomerById(customerId)
	if err != nil {
		log.Println(err.Message)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(err.Code)
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			return
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(customer)
		if err != nil {
			log.Println("getCustomer err: ", err)
		}
	}
}
