package payment

import "fmt"

type CreditCardPayment struct{}

func (creditcardPayment CreditCardPayment) Process() {
	fmt.Println("Payment Process by CrediCard")
}
func init() {
	RegisterPayment("CREDIT_CARD", func() Payment {
		return CreditCardPayment{}
	})
}
