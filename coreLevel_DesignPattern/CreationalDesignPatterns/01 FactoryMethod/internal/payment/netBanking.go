package payment

import "fmt"

type NetBanking struct{}

func (netBanking NetBanking) Process() {
	fmt.Println("Payment done using NetBanking ")
}
func init() {
	RegisterPayment("NET_BANKING", func() Payment {
		return NetBanking{}
	})
}
