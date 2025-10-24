package paytm

import "fmt"

type UPI struct{}

func (u UPI) Process() {
	fmt.Println("Payment process via paytm UPI")
}
