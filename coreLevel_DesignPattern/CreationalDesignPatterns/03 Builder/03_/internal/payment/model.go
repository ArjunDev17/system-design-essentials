package payment

// Payment represents the final complex object built step-by-step
type Payment struct {
	Gateway  string
	Amount   float64
	Currency string
	Method   string
	OrderID  string
	Status   string
}
