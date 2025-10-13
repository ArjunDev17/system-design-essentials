package payment

import "fmt"

// UPIPayment implements Payment interface
type UPIPayment struct{}

func (u UPIPayment) Process() {
	fmt.Println("📱 Processing payment via UPI...")
}

func init() {
	RegisterPayment("UPI", func() Payment {
		return UPIPayment{}
	})
}
