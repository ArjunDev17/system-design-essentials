package payment

import "fmt"

// NetBankingPayment implements Payment interface
type NetBankingPayment struct{}

func (n NetBankingPayment) Process() {
	fmt.Println("🏦 Processing payment via NetBanking...")
}

func init() {
	RegisterPayment("NETBANKING", func() Payment {
		return NetBankingPayment{}
	})
}
