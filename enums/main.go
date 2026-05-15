package main

import "fmt"

// ! By using type keyword we can delcare custom type:

type MyType string
type MyType1 int

// ! Declare type for enum:
type OrderStatus int

// ! Declare enums now:

const (
	Pending OrderStatus = iota
	Received
	Unpaid
	Confirmed
	Cancelled
)

func main() {

	// ! Wrong way to change status: if we receive only string user can pass any string
	// ! We have to define an enum here.
	// func changeStatus(status string) {
	// 	fmt.Println("Order status changed to", status)
	// }

	// ! Mistake can happens. That is why we have to create enums
	// changeStatus("confirmed")
	// changeStatus("confirms")

	// !  Create a function with custom enum type:
	changeOrderStatus := func(status OrderStatus) {
		fmt.Println("You order status is ", status)
	}

	changeOrderStatus(Pending)
	changeOrderStatus(Received)
	changeOrderStatus(Confirmed)
	changeOrderStatus(Cancelled)
	changeOrderStatus(Unpaid)

	// ! Another one example with custom string type:

	// ? Custom user status type
	type UserStatus string

	// ? Declare const block
	const (
		Pending  UserStatus = "pending"
		Approved            = "approved"
		Rejected            = "rejected"
	)

	changeUserStatus := func(status UserStatus) {

		fmt.Println("User status is ", status)

	}

	changeUserStatus(Pending)
	changeUserStatus(Approved)
	changeUserStatus(Rejected)

}

// NOTE: In Go, iota হলো একটা special keyword যেটা mainly const block এর ভিতরে automatic incrementing number generate করার জন্য ব্যবহার করা হয়।
