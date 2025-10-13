package payment

import "fmt"

// CreditCardPayment implements Payment interface
type CreditCardPayment struct{}

func (c CreditCardPayment) Process() {
	fmt.Println("💳 Processing payment via Credit Card...")
}

// Automatically register during package initialization
func init() {
	RegisterPayment("CREDITCARD", func() Payment {
		return CreditCardPayment{}
	})
}
