package payment

import "fmt"

// ProcessCardRazorpay handles Card payments via Razorpay
func ProcessCardRazorpay(amount float64, cardNumber string) {
	fmt.Printf("✅ [Razorpay Card] Payment of ₹%.2f successful for Card: %s\n", 
		amount, maskCardNumber(cardNumber))
	fmt.Println("   → Processed via Razorpay Card Gateway")
}

// Simple card number masking for demo
func maskCardNumber(card string) string {
	if len(card) < 8 {
		return "****"
	}
	return card[:4] + " **** **** " + card[len(card)-4:]
}