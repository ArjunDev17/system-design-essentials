package main

import (
	"fmt"
	"payment-system/internal/factory"
	"payment-system/internal/registry"
)

func main() {
	fmt.Println("🚀 Abstract Factory + Factory Method Example")

	// Register Gateways (Factory Method usage)
	registry.RegisterGateway("RAZORPAY", factory.RazorpayFactory{})
	registry.RegisterGateway("PAYTM", factory.PaytmFactory{})

	// Choose Gateway dynamically (like client choosing gateway)
	fmt.Println("\n--- Using Razorpay Gateway ---")
	razorpayGateway := registry.GetGateway("RAZORPAY")
	razorpayUPI := razorpayGateway.CreateUPI()
	razorpayNB := razorpayGateway.CreateNetBanking()

	razorpayUPI.Process()
	razorpayNB.Process()

	fmt.Println("\n--- Using Paytm Gateway ---")
	paytmGateway := registry.GetGateway("PAYTM")
	paytmUPI := paytmGateway.CreateUPI()
	paytmNB := paytmGateway.CreateNetBanking()

	paytmUPI.Process()
	paytmNB.Process()
}
