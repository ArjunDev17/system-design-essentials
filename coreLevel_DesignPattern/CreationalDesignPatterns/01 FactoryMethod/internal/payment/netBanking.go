package payment

import "fmt"

type NetBanking struct{}

func (netBanking NetBanking) Process() {
	fmt.Println("Payment done using NetBanking ")
}
