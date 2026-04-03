package payment

import "fmt"

// ProcessCardStripe handles Card payments via Stripe
func ProcessCardStripe(amount float64, cardNumber string) {
	fmt.Printf("✅ [Stripe Card] Payment of ₹%.2f successful for Card: %s\n",
		amount, maskCardNumber(cardNumber))
	fmt.Println("   → Processed via Stripe Card Gateway")
}
