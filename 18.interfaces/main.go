package main

import "fmt"

type paymenter interface {
	pay(amount int)
}

// ! Create a struct for payment implemention
// type payment struct {
// 	razorpay razorpay
// 	stripe stripe
// }

type payment struct {
	// gateway stripe
	gateway paymenter
}

// ! Create make payment function for payment :
func (p payment) makePayment(amount int) {
	// p.razorpay.pay(amount)
	p.gateway.pay(amount)
}

// ! Create another struct to pay with razor pay:
type razorpay struct {
}

// ! add  an function with razor pay:
func (r razorpay) pay(amount int) {
	fmt.Println(`Amount has been paid  by using razor pay`, amount)

}

// Configure stripe struct
type stripe struct {
}

func (s stripe) pay(amount int) {
	fmt.Println("Amount has been paid using stripe", amount)
}

// Type fack payment gateway:

type fakeGateway struct {
}

func (f fakeGateway) pay(amount int) {
	fmt.Println("Now making payment with fake gateway", amount)
}

func main() {

	// ! create an instance for razorpay first:
	// razorpayPaymentSettings := razorpay{}

	// ! Create an interface for stripe:
	// stripePaymentSettings := stripe{}

	// ! Create an instance for fakeGateway
	fakeGatewayPaymentSettings := fakeGateway{}

	// ! Create one instance  for payment
	paymentSystem := payment{
		// razorpay: razorpayPaymentSettings,
		// stripe: stripePaymentSettings,
		// gateway: stripePaymentSettings,
		gateway: fakeGatewayPaymentSettings,
	}

	paymentSystem.makePayment(200)

}
