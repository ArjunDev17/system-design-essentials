package paytm

import "fmt"

type NetBanking struct{}

func (nb NetBanking) Process() {
	fmt.Println("Paytm Payment Process done by netbanking")
}
