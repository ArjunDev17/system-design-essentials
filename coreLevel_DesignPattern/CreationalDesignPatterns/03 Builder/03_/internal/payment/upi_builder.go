package payment

// UPIPaymentBuilder is a concrete builder for UPI payments
type UPIPaymentBuilder struct {
	payment *Payment
}

// NewUPIPaymentBuilder initializes the UPI builder
func NewUPIPaymentBuilder() *UPIPaymentBuilder {
	return &UPIPaymentBuilder{payment: &Payment{}}
}

func (b *UPIPaymentBuilder) SetGateway() {
	b.payment.Gateway = "Razorpay-UPI"
}

func (b *UPIPaymentBuilder) SetAmount(amount float64) {
	b.payment.Amount = amount
}

func (b *UPIPaymentBuilder) SetCurrency(currency string) {
	b.payment.Currency = currency
}

func (b *UPIPaymentBuilder) SetMethod() {
	b.payment.Method = "UPI"
}

func (b *UPIPaymentBuilder) SetOrderID(orderID string) {
	b.payment.OrderID = orderID
}

func (b *UPIPaymentBuilder) SetStatus() {
	b.payment.Status = "Pending"
}

func (b *UPIPaymentBuilder) GetPayment() *Payment {
	return b.payment
}
