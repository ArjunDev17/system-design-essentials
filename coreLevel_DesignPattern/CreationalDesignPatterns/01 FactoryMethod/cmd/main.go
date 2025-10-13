package main

import (
	"fmt"
	"payment-factory-pattern/internal/payment"
)

func main() {
	fmt.Println("=== OCP-Compliant Payment Factory Example ===")

	p1 := payment.GetPaymentMethod("UPI")
	p1.Process()

	p2 := payment.GetPaymentMethod("CREDITCARD")
	p2.Process()

	p3 := payment.GetPaymentMethod("NETBANKING")
	p3.Process()
}
