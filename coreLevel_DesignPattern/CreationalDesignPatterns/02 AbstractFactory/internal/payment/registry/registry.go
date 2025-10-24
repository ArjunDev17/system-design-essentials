package registry

import "payment-system/internal/factory"

// gatewayRegistry holds all registered payment gateways.
var gatewayRegistry = make(map[string]factory.PaymentGatewayFactory)

// RegisterGateway registers a payment gateway.
func RegisterGateway(name string, gateway factory.PaymentGatewayFactory) {
	gatewayRegistry[name] = gateway
}

// GetGateway retrieves a registered payment gateway.
func GetGateway(name string) factory.PaymentGatewayFactory {
	g, exists := gatewayRegistry[name]
	if !exists {
		panic("unknown payment gateway: " + name)
	}
	return g
}
