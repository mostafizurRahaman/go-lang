package main

import "fmt"

// ! declaration of enum:
type OrderStatus int

// ! Declaration :

const (
	PENDING OrderStatus = iota
	CONFIRMATION
	REJECTION
)

type UserStatus string

// ! Change user status :
const (
	Pending     UserStatus = "pending"
	OtpVerified            = "otp_verified"
	Active                 = "active"
	InReview               = "in_review"
	Blocked                = "blocked"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status is ", status)
}

func changeUserStatus(status UserStatus) {
	fmt.Println("User status changed to", status)
}

func main() {

	// ! Change order status change:
	changeOrderStatus(PENDING)
	changeOrderStatus(CONFIRMATION)
	changeOrderStatus(REJECTION)

	// ! Change User Status:
	changeUserStatus(Pending)
	changeUserStatus(Active)
	changeUserStatus(OtpVerified)
	changeUserStatus(InReview)
	changeUserStatus(Blocked)

}
