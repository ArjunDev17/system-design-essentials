package main

import (
	"fmt"

	"builder-pattern/internal/payment"
)

func main() {
	// Director using UPI Builder
	upiBuilder := payment.NewUPIPaymentBuilder()
	director := payment.NewPaymentDirector(upiBuilder)
	upiPayment := director.BuildPayment(5000, "INR", "ORD_001")

	// Director using Card Builder
	cardBuilder := payment.NewCardPaymentBuilder()
	director = payment.NewPaymentDirector(cardBuilder)
	cardPayment := director.BuildPayment(12000, "USD", "ORD_002")

	fmt.Printf("✅ UPI Payment: %+v\n", upiPayment)
	fmt.Printf("💳 Card Payment: %+v\n", cardPayment)
}
