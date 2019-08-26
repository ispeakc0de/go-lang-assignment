package domain

//declare the customer structure
type Customer struct{
ID,Name,Email string
}
//interface to declare all the the functions
type CustomerStore interface {
Create (Customer) error
Update (string, Customer) error
Delete(string) error
GetById(string) (Customer, error)
GetAll() ([]Customer, error)
}
