package payment

import "fmt"

type UPIPayment struct{}

func (up UPIPayment) Process() {
	fmt.Println("Payment Done By UPI")
}
