package mapstore
import (
"errors"
"fmt"
"assign2/domain"
)

//Mapstore  structure

type MapStore struct{
store map[string]domain.Customer
}

func NewMapStore() *MapStore {
 return &MapStore { store: make(map[string]domain.Customer)}
 }

 func (ms *MapStore) Create(c domain.Customer) error {
	if c.ID == "" {
		return errors.New("ID not found for the customer")
	}
	ms.store[c.ID] = c
	fmt.Println("User created successfully, User ID:", c.ID)
	return nil
}

func (ms *MapStore) Delete(ID string) error {
	_, ok := ms.store[ID]
	if ok {
		delete(ms.store, ID)
		fmt.Println("User deleted successfully")
		return nil
	}
	return errors.New("Cannot delete: the customer does not exist")
}

func (ms *MapStore) Update(ID string, c domain.Customer) error {
_, ok := ms.store[ID]
	if ok {
		ms.store[ID] = c
		fmt.Println("updated successfully")
		return nil
	}
	return errors.New("Cannot update : the customer does not exist")

}

func (ms *MapStore) GetById(ID string) (domain.Customer, error) {

	val, ok := ms.store[ID]
	if ok {
		return val,nil
	}
	return domain.Customer{},errors.New("the customer does not exist")

	// customers, err := ms.GetAll()

}

func (ms *MapStore) GetAll() ([]domain.Customer, error) {
var cust [] domain.Customer
if len(ms.store)!=0 {
	for _, v := range ms.store{
		cust = append(cust,v)
	}
	return cust,nil
	}

	return nil,errors.New("No record found!")
}
