package payment

import (
	"fmt"
)

// registry holds all registered payment types
var registry = make(map[string]func() Payment)

// RegisterPayment registers a new payment type in the factory
func RegisterPayment(name string, constructor func() Payment) {
	registry[name] = constructor
}

// GetPaymentMethod retrieves the payment implementation
func GetPaymentMethod(method string) Payment {
	constructor, exists := registry[method]
	if !exists {
		panic(fmt.Sprintf("❌ Unknown payment method: %s", method))
	}
	return constructor()
}
