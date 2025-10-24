package razorpay

import "fmt"

type UPI struct{}

func (upi UPI) Process() {
	fmt.Println("RZP payment done by upi")
}
