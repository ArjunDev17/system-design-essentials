package factory

import "payment-sys/internal/methods"

type PaymentGateWayFactory interface {
	CreateUPI() methods.Payment
	CreateNetBanking() methods.Payment
}
