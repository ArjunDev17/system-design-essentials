package main

import (
	"fmt"
	"payment-factory-pattern/internal/payment"
)

func main() {
	fmt.Println("factory pattern method example")

	p1 := payment.GetPaymentWay("UPI")
	p1.Process()

	p2 := payment.GetPaymentWay("CREDIT_CARD")
	p2.Process()

	p3 := payment.GetPaymentWay("NET_BANKING")
	p3.Process()
}
