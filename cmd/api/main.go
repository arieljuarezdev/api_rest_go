package main

import (
	"net/http"

	ctrls "api_rest/internal/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//listagem de clientes
	r.HandleFunc("/customers", ctrls.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", ctrls.GetCustumerById).Methods("GET")
	r.HandleFunc("/newCustomer", ctrls.InsertCustumer).Methods("POST")
	r.HandleFunc("/updateCustomer/{id}", ctrls.UpdateCustomer).Methods("PATCH")
	r.HandleFunc("/deleteCustomer/{id}", ctrls.DeleteCustomer).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
