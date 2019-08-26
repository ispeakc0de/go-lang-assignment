package main

import (
	"fmt"
	"assign2/domain"
	"assign2/mapstore"
	"testing"
    . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type CustomerController struct {
	Store domain.CustomerStore
}

var controller = CustomerController{
	Store: mapstore.NewMapStore(),
}
var customer = domain.Customer{
				ID:   "1",
				Name: "Shubham Chaudhary",
				Email: "ashubam314@gmail.com",
			}
var custUpdate = domain.Customer{
		        ID:   "1",
		        Name: "Shubham",
		        Email: "ashub4@gmail.com",
			}

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = Describe("BDD on Customer App", func() {

	Context("Customer add function testing", func() {

		It("It should return a nil value", func() {

			err := controller.Store.Create(customer)
			if err != nil {
				fmt.Println("error", err)
			}
			Expect(err).To(BeNil())

		})
	})

	Context("Customer update function testing", func() {

		It("It should return a nil Value", func() {

			err := controller.Store.Update(custUpdate.ID, custUpdate)
			if err != nil {
				fmt.Println("error", err)
			}

			Expect(err).To(BeNil())

		})
	})

	Context("Customer GetByID function testing", func() {

		It("It should return a nil value", func() {

			_, err := controller.Store.GetById("1")
			Expect(err).To(BeNil())

		})
	})

	Context("Customer GetAll function testing", func() {

		It(" It should return a nil value", func() {

			_, err := controller.Store.GetAll()

			Expect(err).To(BeNil())

		})
	})

	Context("Customer delete function testing", func() {

		It("It should return a  nil value", func() {

			err := controller.Store.Delete("1")
			Expect(err).To(BeNil())

		})
	})

})
