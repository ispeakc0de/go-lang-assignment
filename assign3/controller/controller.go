package controller
import (
	"encoding/json"
	"fmt"
	"assign3/domain"
	"assign3/mapstore"
	"net/http"
    "mux"
)

// customerController structure
type CustomerController struct {
Store domain.CustomerStore
}

// creating object of type customerController
var controller = CustomerController{
	Store: mapstore.NewMapStore(),
}

// add a value to the mapstore which is provided in url at port number 8080.
func Add(w http.ResponseWriter, r *http.Request) {
      vars := mux.Vars(r)
	customer := domain.Customer{
		ID:  vars["id"],
		Name: vars["name"],
		Email: vars["email"],
	}

	err := controller.Store.Create(customer)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result, err := json.Marshal("user created sucessfully ")
    w.Write(result)

}

//update the record in mapstore to a given Id which is provided in url at port number 8080.
func Update(w http.ResponseWriter, r *http.Request) {
 vars := mux.Vars(r)
 var pid=vars["pid"]
	updateCustomer := domain.Customer{
		ID:   vars["id"],
		Name: vars["name"],
		Email: vars["email"],
	}
	err := controller.Store.Update(pid, updateCustomer)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("updated:", updateCustomer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result, err := json.Marshal("user updated sucessfully")
    w.Write(result)

}

//delete the record from the mapstore to a give key which is provided in url at port number 8080.
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := controller.Store.Delete(vars["id"])
	if result != nil {
		fmt.Println("error", result)
		return
	}
	fmt.Fprintf(w, "record Deleted ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rs,_:= json.Marshal(result)
	w.Write(rs)
}

//return the record with given key ,which is provided in url at port number 8080.
func GetByID(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	result, err := controller.Store.GetById(vars["id"])
	if err != nil {
		fmt.Println("error", err)
		rs,_ := json.Marshal("No record found")
	    w.Write(rs)
		return
	}
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rs, err := json.Marshal(result)
	w.Write(rs)
}


//Return all records from the mapstore
func GetAll(w http.ResponseWriter, r *http.Request) {
    result, err := controller.Store.GetAll()
	if err != nil {
		fmt.Println("error ",err)
		rs,_ := json.Marshal(" No record found")
	    w.Write(rs)
		return
	}
	fmt.Println("GetAll:", result)
	fmt.Fprintf(w, "GetAll", result)
	rs, err := json.Marshal(result)
	w.Write(rs)
}