package factory

import "payment-system/internal/methods"

type PaymentGatewayFactory interface {
	CreateUPI() methods.Payment
	CreateNetBanking() methods.Payment
}
