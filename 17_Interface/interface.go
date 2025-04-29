package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gatway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpaygw := razorpay{}
	// razorpaygw.pay(amount)

	// paytmgw := paytm{}
	// paytmgw.pay(amount)
	p.gatway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using Razorpay")
}

// type paytm struct{}

// func (p paytm) pay(amount float32) {
// 	fmt.Println("Payment of", amount, "made using Paytm")
// }

type fake struct{}

func (f fake) pay(amount float32) {
	fmt.Println("Payment of", amount, "made using Fake")
}

func main() {

	razorpaygw := razorpay{}

	newpayment := payment{
		gatway: razorpaygw,
	}
	newpayment.makePayment(1000.0)

}
