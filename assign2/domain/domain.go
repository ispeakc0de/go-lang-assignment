package domain
//structure
type Customer struct{
ID,Name,Email string
}
//interface
type CustomerStore interface {
Create (Customer) error
Update (string, Customer) error
Delete(string) error
GetById(string) (Customer, error)
GetAll() ([]Customer, error)
}
