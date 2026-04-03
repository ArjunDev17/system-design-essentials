package payment

import "fmt"

// ProcessUPI routes UPI payment to the correct provider
func ProcessUPI(provider string, amount float64, upiID string) {
	switch provider {
	case "razorpay":
		ProcessUPIRazorpay(amount, upiID)
	case "stripe":
		ProcessUPIStripe(amount, upiID)
	default:
		fmt.Printf("❌ Unknown UPI provider: %s\n", provider)
	}
}

// ProcessCard routes Card payment to the correct provider
func ProcessCard(provider string, amount float64, cardNumber string) {
	switch provider {
	case "razorpay":
		ProcessCardRazorpay(amount, cardNumber)
	case "stripe":
		ProcessCardStripe(amount, cardNumber)
	default:
		fmt.Printf("❌ Unknown Card provider: %s\n", provider)
	}
}