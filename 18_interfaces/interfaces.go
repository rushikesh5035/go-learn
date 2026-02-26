package main

import "fmt"

// Interface
// here we have two payment gateway razorpay and stripe, both of them have same method pay but different implementation.
// so inside the interface we gives method
type paymenter interface {
	pay(amount float32)
	
}


type payment struct {
	// gateway stripe
	// gateway razorpay

	// after creating interface now our gateway is paymeter type
	gateway paymenter
}

// open closed principle: open for extension but closed for modification (we can add new payment gateway without modifying existing code)
// in the above code, if we want to add new payment gateway, we need to modify the makePayment method of payment struct which is not a good practice. we can use interface to achieve this.

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)

	p.gateway.pay(amount)
}

type razorpay struct {}
func (r razorpay) pay(amount float32) {
	// logic to make payment using razorpay
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct {}
func (s stripe) pay(amount float32) {
	// logic to make payment using stripe
	fmt.Println("Making payment using stripe", amount)
}

type fakePayment struct {}
func (f fakePayment) pay(amount float32) {
	fmt.Println("Making payment using fake gateway for testing", amount)
}

type paypal struct {}
func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using paypal", amount)
}

func main() {
	// instance of payment
	// payment := payment{}
	// payment.makePayment(100.0)

	// stripe:= stripe{}
	// stripe.pay(200.0)

	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	// fakePaymentGw := fakePayment{}
	paypalPaymentGw := paypal{}

	
	newPayment := payment{
		// gateway: stripePaymentGw,
		// gateway: razorpayPaymentGw, // it give us error because we need to change the gateway from stripe to razorpay in the payment struct 

		// after creating interface we can easily change the gateway without modifying the payment struct
		gateway: paypalPaymentGw, // gives same error as above
	}

	newPayment.makePayment(300.0)

}