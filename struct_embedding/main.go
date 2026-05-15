package main

import (
	"fmt"
	"time"
)

type customer struct {
	name      string
	email     string
	createdAt time.Time
}

func (c customer) getCustomerName() string {
	return c.name
}

func (c *customer) updateCustomerName(name string) {

	c.name = name
}

type order struct {
	productId   int
	productName string
	customer    customer
}

func main() {

	// ! Create an instance for customer struct:
	customer1 := customer{
		name:      "Tamim hossain",
		email:     "tamim@gamil.com",
		createdAt: time.Now(),
	}
	fmt.Println(customer1)

	// ! Create an instance from order struct:
	order1 := order{
		productName: "New Product",
		productId:   2,
		customer: customer{
			name:      "Rahim hossain",
			email:     "rahim@gmail.com",
			createdAt: time.Now(),
		},
	}

	fmt.Println(order1)

	// ! Another way create instance for embaded struct:
	order2 := order{
		productName: "New order 2",
		productId:   3,
		customer:    customer1,
	}

	order2.customer.email = "rahim2343@gmail.com"

	fmt.Println(order2)

	// ! access customer functions by order:

	order2.customer.updateCustomerName("Rahim Hossain Updatead Name")

	fmt.Println(order2.customer.getCustomerName())

	fmt.Println(order2)
}
