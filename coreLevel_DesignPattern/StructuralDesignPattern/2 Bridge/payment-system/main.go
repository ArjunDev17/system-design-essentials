package main

import (
	"fmt"
	"payment-system/payment"
)

func main() {
	fmt.Println("=== Payment System Demo ===\n")

	// UPI Payments
	fmt.Println("--- UPI Payments ---")
	payment.ProcessUPI("razorpay", 500.00, "user@upi")
	payment.ProcessUPI("stripe", 750.00, "user@upi")

	// Card Payments
	fmt.Println("\n--- Card Payments ---")
	payment.ProcessCard("razorpay", 1200.00, "4111 1111 1111 1111")
	payment.ProcessCard("stripe", 2500.00, "4242 4242 4242 4242")

	fmt.Println("\n✅ Payment system executed successfully!")
}