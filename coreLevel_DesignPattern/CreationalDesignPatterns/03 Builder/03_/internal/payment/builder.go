package payment

// PaymentBuilder defines the interface for building a Payment step by step
type PaymentBuilder interface {
	SetGateway()
	SetAmount(amount float64)
	SetCurrency(currency string)
	SetMethod()
	SetOrderID(orderID string)
	SetStatus()
	GetPayment() *Payment
}
