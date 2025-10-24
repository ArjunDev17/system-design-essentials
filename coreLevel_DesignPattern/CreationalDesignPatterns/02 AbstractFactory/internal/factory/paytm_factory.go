package factory

import (
	"payment-system/internal/methods"
	"payment-system/internal/paytm"
)

// PaytmFactory creates Paytm payment methods.
type PaytmFactory struct{}

func (p PaytmFactory) CreateUPI() methods.Payment {
	return paytm.UPI{}
}

func (p PaytmFactory) CreateNetBanking() methods.Payment {
	return paytm.NetBanking{}
}
