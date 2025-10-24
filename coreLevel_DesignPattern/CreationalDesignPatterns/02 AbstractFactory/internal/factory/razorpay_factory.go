package factory

import (
	"payment-system/internal/methods"
	"payment-system/internal/payment/razorpay"
)

type RazorpayFactory struct{}

func (r RazorpayFactory) CreateUPI() methods.Payment {
	return razorpay.UPI{}
}

func (r RazorpayFactory) CreateNetBanking() methods.Payment {
	return razorpay.NetBanking{}
}
