package main
import (
"fmt"
"assign2/domain"
"assign2/mapstore"
)
type CustomerController struct {
store domain.CustomerStore
}
//use create function
  func (C CustomerController ) Add (c domain.Customer) {
    err:= C.store.Create(c)
    if err!=nil {
    fmt.Println("Error:", err)
    return
     }
     fmt.Println("New Customer has been created")
     }

//use delete function

func (C CustomerController ) DeleteValue (c domain.Customer) {
    err:= C.store.Delete(c.ID)
    if err!=nil {
    fmt.Println("Error:", err)
    return
     }
     fmt.Println("Customer deleted")
     }

//print

func (C CustomerController ) PrintAll () {
    val,err := C.store.GetAll()
    if err!=nil {
    fmt.Println("Error:", err)
    return
     }
     for _,v := range val{
		fmt.Println(v)
	}
     }

     //print by ID

func (C CustomerController ) PrintById (id string) {
    val,err := C.store.GetById(id)
    if err!=nil {
    fmt.Println("Error:", err)
    return
     }
		fmt.Println(val)

     }

     //update customer

     func (C CustomerController ) Update (id string,c domain.Customer) {
    err := C.store.Update(id,c)
    if err!=nil {
    fmt.Println("Error:", err)
    return
     }
		fmt.Println("UPDATED SUCCESSFULLY!")

     }

  func main() {
 controller := CustomerController{
  store : mapstore.NewMapStore(),
  //store: = mongodb.NewMongoStore()
  }
  customer := domain.Customer {
  ID : "cust101",
   Name: "JPM",
     }
      customer2 := domain.Customer {
  ID : "cust102",
   Name: "shubh",
     }

     controller.Add(customer)
      controller.Add(customer2)
     controller.PrintAll()
       controller.PrintById("cust101")
     //controller.DeleteValue(customer)
     controller.Update("cust101",customer2)
     controller.PrintAll()
     controller.Add(customer)
     controller.PrintAll()

     }


