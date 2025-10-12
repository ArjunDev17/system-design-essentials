package payment

import "fmt"

type CreditCardPayment struct{}

func (creditcardPayment CreditCardPayment) Process() {
	fmt.Println("Payment Process by CrediCard")
}
