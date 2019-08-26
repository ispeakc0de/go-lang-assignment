package main

import (
	"fmt"
	"assign3/controller"
	"net/http"
	"net/http/httptest"
	"testing"
    "mux"
)

//It will test add function.Test case is only pass when it response with a status code 200
func TestAdd(t *testing.T) {
// create a instance of router
	r := mux.NewRouter()

    //declare the handler for add function
	r.HandleFunc("/add/{id}/{name}/{email}", controller.Add).Methods("GET")

	//http request for add record
	req, err := http.NewRequest("GET","/add/1/shubh/ashu@gmail", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	// error code 200 means the test case is pass
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}


//It will test Update function.Test case is only pass when it response with a status code 200
func TestUpdate(t *testing.T) {
	r := mux.NewRouter()
//define a handler function for update
	r.HandleFunc("/update/{pid}/{id}/{name}/{email}", controller.Update).Methods("GET")

	//http request for update
	req, err := http.NewRequest("GET", "/update/1/1/shubh/ash@gmail", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}


//It will test GetById function.Test case is only pass when it response with a status code 200
func TestGetByID(t *testing.T) {
	r := mux.NewRouter()
//define handler function for get method
	r.HandleFunc("/get/{id}", controller.GetByID).Methods("GET")

	// http request for get
	req, err := http.NewRequest("GET", "/get/1", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}


//It will test delete function.Test case is only pass when it response with a status code 200
func TestDelete(t *testing.T) {
	r := mux.NewRouter()

//define handler function for delete method
	r.HandleFunc("/delete/{id}", controller.Delete).Methods("GET")
	//http request for delete
	req, err := http.NewRequest("GET", "/delete/1", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}


//It will test getall function.Test case is only pass when it response with a status code 200
func TestGetAll(t *testing.T) {
	r := mux.NewRouter()
//define handler function for getall
	r.HandleFunc("/getall", controller.GetAll).Methods("GET")
	//http request for getall
	req, err := http.NewRequest("GET", "/getall", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}
}