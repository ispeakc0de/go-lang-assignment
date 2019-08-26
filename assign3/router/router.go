package router
import (
"mux"
"assign3/controller"
)
//create an instance of router
var Mux=mux.NewRouter()

//declare all the handlers
func Routers () {
	Mux.HandleFunc("/add/{id}/{name}/{email}", controller.Add).Methods("GET")
	Mux.HandleFunc("/update/{pid}/{id}/{name}/{email}", controller.Update) .Methods("GET")
	Mux.HandleFunc("/get/{id}", controller.GetByID).Methods("GET")
	Mux.HandleFunc("/getall", controller.GetAll).Methods("GET")
	Mux.HandleFunc("/delete/{id}", controller.Delete).Methods("GET")

}
