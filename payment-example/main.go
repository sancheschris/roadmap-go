package main

import "payment-example/payment"

func Checkout(processor payment.PaymentProcessor, amount float64) {
	processor.Pay(amount)
}

func main() {
	paypal := payment.PaypalProcessor{}
	stripe := payment.StripeProcessor{}

	Checkout(paypal, 100.0)
	Checkout(stripe, 200.0)
}


