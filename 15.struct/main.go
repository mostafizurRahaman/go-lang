package main

import (
	"fmt"
	"time"
)

// NOTE: What we have learnt :
// NOTE: - 1. Struct declaration
// Note: - 2. Inline struct
// Note: - 3. Create struct instance
// Note: - 4. update and add instance properties values
// Note: - 5. Access and update instance properties from outside.
// Note: - 6. Create a constructor for struct
// Note: - 7. Create getter and setter function for instance
// Note: - 8. If we don't assign a value to instance any properties it will be zero value ("",0,false, nil):

// Type :
// ! Build a struct
type order struct {
	id        int
	name      string
	email     string
	age       int
	status    string
	createdAt time.Time
	updatedAt time.Time
}

// ! Get email function :
func (o order) getEmail() string {

	return o.email
}

// ! To update a value create a method which will update the value:
//
//	(o *order) this is called receiver type:
func (o *order) updateStatus(status string) {
	o.status = status
}

// ! Go don't have class and constructure. So for struct here we have creating an constructor function which will recieve all the required properties as parameter and return that struct instance:
// ** Here we have to return pointer value

func newOrder(id, age int, name, email, status string) *order {

	myOrder := order{
		id:        id,
		name:      name,
		email:     email,
		status:    status,
		age:       age,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	return &myOrder
}

func main() {

	// ! First way to create struct instance
	var person order
	fmt.Println(person) // Empty struct whill always contain zero value for each field.
	// fmt.Println(person == nil)
	person = order{
		id:        1,
		name:      "Mostafizur Rahaman",
		age:       24,
		email:     "mostafizur@gmail.com",
		status:    "pending",
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	fmt.Println(person)

	// ! Second way to create a struct instance:

	person2 := order{
		id:        2,
		name:      "Azmir Azam khan",
		age:       25,
		email:     "azmir@gmail.com",
		status:    "confirmed",
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	fmt.Println(person2)

	// ! Access struct values:
	fmt.Println(person2.age)
	fmt.Println(person.age)

	// ! Change instance any properties value:
	person2.email = "azmirazamkhan@gmail.com"
	fmt.Println(person2.email)
	fmt.Println(len(person2.name))

	// ! Call function to get email:
	fmt.Println(person2.getEmail())
	fmt.Println(person.getEmail())

	// ! Call function to update the status:
	person.updateStatus("confirmed")
	person2.updateStatus("rejected")
	fmt.Println(person)
	fmt.Println(person2)

	person3 := newOrder(3, 20, "Mostafizur", "fh@gmail.com", "confirmed")

	fmt.Println(person3)
	person3.updateStatus("rejected")

	// !Declare struct direclty and initialize value while declaring
	// !Inline struct
	language := struct {
		name  string
		email string
	}{
		name:  "Mostafizur",
		email: "mostafizurrahaman0401@gmail.com",
	}
	fmt.Println(language)
}
