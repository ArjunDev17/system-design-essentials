package razorpay

import "fmt"

type NetBanking struct{}

func (nb NetBanking) Process() {
	fmt.Println("RZP payment done  by netbanking ")
}
