package main

import (
	"fmt"
	"time"
)

// !1. Declare an struct:
type order struct {
	id        int
	name      string
	email     string
	status    string
	createdAt time.Time
	updatedAt time.Time
}

// ! Create getter function :
func (o order) getOrderEmailAndStatus() (string, string) {
	return o.email, o.name
}

// ! create a setter function :
func (o *order) setOrderStatus(status string) {

	o.status = status
}

// ! Create a constructor function for this class:

func newOrder(id int, name string, email string, status string) *order {

	myOrder := order{
		id:        id,
		name:      name,
		email:     email,
		status:    status,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	return &myOrder

}

func main() {

	// ! 1.  Create an instance from struct: (Only declaration without initialization)
	var myOrder order
	fmt.Println(myOrder) // it will nil because I don't initialize the value:

	// ! 2. Declare and initiailiation:
	yourOrder := order{
		id:        1,
		name:      "Osman goni",
		email:     "goni@gmail.com",
		status:    "pending",
		createdAt: time.Now(), // time in nanosecond fraction.
		updatedAt: time.Now(),
	}

	fmt.Println(yourOrder)

	// ! get Properties values : with dot notation
	fmt.Println(yourOrder.email)

	// ! set properties values with dot notion :
	yourOrder.email = "osman@gmail.com"

	// ! if any properties is not assigned the properties initial values will be falsy value/ zero values.

	// ! Call the getter function to get properties values :
	email, name := yourOrder.getOrderEmailAndStatus()
	fmt.Println(email, name)

	// ! Call the setter function to set the value:
	yourOrder.setOrderStatus("confirmed")
	fmt.Println(yourOrder)

	// another one instance:
	secondOrder := order{
		id:        3,
		name:      "Rafi Hasan",
		email:     "rafi@gmail.com",
		status:    "pending",
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	fmt.Println(secondOrder)

	// ! Create another instance from constructor:
	thirdOrder := newOrder(4, "Nm sujon", "nm@gmail.com", "pending")
	fmt.Println(thirdOrder)

	// ! Create an  struct for single instance only:
	// ! declaration and instance  creation  same line

	language := struct {
		secret_key string
		web_secret string
	}{
		secret_key: "Mostafiz",
		web_secret: "Ratul Hossain",
	}

	fmt.Println(language)
	fmt.Println(language.secret_key)
}
