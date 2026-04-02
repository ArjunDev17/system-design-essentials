package payment

import "fmt"

// ProcessUPIStripe handles UPI payments via Stripe
func ProcessUPIStripe(amount float64, upiID string) {
	fmt.Printf("✅ [Stripe UPI] Payment of ₹%.2f successful for UPI ID: %s\n", amount, upiID)
	fmt.Println("   → Processed via Stripe UPI (India)")
}