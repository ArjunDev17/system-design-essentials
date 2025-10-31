package payment

// CardPaymentBuilder is a concrete builder for Card payments
type CardPaymentBuilder struct {
	payment *Payment
}

// NewCardPaymentBuilder initializes the Card builder
func NewCardPaymentBuilder() *CardPaymentBuilder {
	return &CardPaymentBuilder{payment: &Payment{}}
}

func (b *CardPaymentBuilder) SetGateway() {
	b.payment.Gateway = "Razorpay-Card"
}

func (b *CardPaymentBuilder) SetAmount(amount float64) {
	b.payment.Amount = amount
}

func (b *CardPaymentBuilder) SetCurrency(currency string) {
	b.payment.Currency = currency
}

func (b *CardPaymentBuilder) SetMethod() {
	b.payment.Method = "Card"
}

func (b *CardPaymentBuilder) SetOrderID(orderID string) {
	b.payment.OrderID = orderID
}

func (b *CardPaymentBuilder) SetStatus() {
	b.payment.Status = "Initialized"
}

func (b *CardPaymentBuilder) GetPayment() *Payment {
	return b.payment
}
