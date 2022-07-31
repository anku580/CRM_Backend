package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var database = map[string]Customer{}
var uniqueCustId = 3

func getAllCustomers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(database)
}

func getCustomerById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customerId, _ := mux.Vars(r)["id"]
	found := false

	for k, v := range database {
		fmt.Println(v)
		if k == customerId {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(v)
			found = true
		}
	}

	if found == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No member Found!")
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	newEntry := new(Customer)
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)
	uniqueCustId = uniqueCustId + 1
	newEntry.Id = strconv.Itoa(uniqueCustId)
	if _, ok := database[newEntry.Id]; ok {
		//conflict
		w.WriteHeader(http.StatusConflict)
	} else {
		database[newEntry.Id] = *newEntry
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(database)

}

func updateCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	customerId, _ := mux.Vars(r)["id"]
	found := false

	newEntry := new(Customer)

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)
	newEntry.Id = customerId
	for k, v := range database {
		fmt.Println(v)
		if k == customerId {
			database[customerId] = *newEntry
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(database)
			found = true
		}
	}

	if found == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No member Found!")
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customerId, _ := mux.Vars(r)["id"]

	if _, ok := database[customerId]; ok {
		delete(database, customerId)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(database)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(database)
	}
}

func main() {

	cust := Customer{Id: "1", Name: "Anku", Role: "admin", Email: "anku580@gmail.com", Phone: 1234, Contacted: true}
	cust2 := Customer{Id: "2", Name: "John", Role: "user", Email: "john@gmail.com", Phone: 5673444, Contacted: false}
	cust3 := Customer{Id: "3", Name: "Jacob", Role: "user", Email: "jacob@gmail.com", Phone: 56734414, Contacted: false}

	database["1"] = cust
	database["2"] = cust2
	database["3"] = cust3

	router := mux.NewRouter()

	router.HandleFunc("/getAllCustomers", getAllCustomers).Methods("GET")
	router.HandleFunc("/getCustomer/{id}", getCustomerById).Methods("GET")
	router.HandleFunc("/createCustomer", createCustomer).Methods("POST")
	router.HandleFunc("/updateCustomer/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/deleteCustomer/{id}", deleteCustomer).Methods("DELETE")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
