package payment

import "fmt"


type PaymentProcessor interface {
	Pay(amount float64) error
}

type PaypalProcessor struct{}

func (p PaypalProcessor) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Paypal\n", amount)
	return nil
}

type StripeProcessor struct{}

func (p StripeProcessor) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Stripe\n", amount)
	return nil
}