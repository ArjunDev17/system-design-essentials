package payment

// PaymentDirector controls the building process
type PaymentDirector struct {
	builder PaymentBuilder
}

// NewPaymentDirector initializes a new director
func NewPaymentDirector(b PaymentBuilder) *PaymentDirector {
	return &PaymentDirector{builder: b}
}

// BuildPayment orchestrates the building steps
func (d *PaymentDirector) BuildPayment(amount float64, currency, orderID string) *Payment {
	d.builder.SetGateway()
	d.builder.SetAmount(amount)
	d.builder.SetCurrency(currency)
	d.builder.SetMethod()
	d.builder.SetOrderID(orderID)
	d.builder.SetStatus()
	return d.builder.GetPayment()
}
