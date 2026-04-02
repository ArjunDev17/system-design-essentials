package payment

import "fmt"

// ProcessUPIRazorpay handles UPI payments via Razorpay
func ProcessUPIRazorpay(amount float64, upiID string) {
	fmt.Printf("✅ [Razorpay UPI] Payment of ₹%.2f successful for UPI ID: %s\n", amount, upiID)
	fmt.Println("   → Processed via Razorpay UPI Gateway")
}