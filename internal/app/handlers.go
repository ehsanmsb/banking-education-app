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
	//customers := []customer{
	//	{Name: "Ehsan", Email: "ehsan.mosayebi@snapp.cab", City: "Tehran", ZipCode: "123456"},
	//	{Name: "Vahid", Email: "vahid.fardi@snapp.cab", City: "Tehran", ZipCode: "456788"},
	//	{Name: "Mostafa", Email: "mostafa.negim@snapp.cab", City: "Tehran", ZipCode: "8453303"},
	//	{Name: "Reza", Email: "reza.alizadeh@snapp.cab", City: "Tehran", ZipCode: "8906453"},
	//}
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
		log.Println(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(customer)
		if err != nil {
			log.Println("getCustomer err: ", err)
		}
	}
}
