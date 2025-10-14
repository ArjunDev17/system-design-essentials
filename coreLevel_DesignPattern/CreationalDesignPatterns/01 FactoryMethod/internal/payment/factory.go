package payment


var registry=make(map[string]func()Payment)

func RegisterPayment(name string,costructor func() Payment)  {
	registry[name]=costructor
}

func GetPaymentWay(way string) Payment {
	constructor,exist:=registry[way]
	if !exist{
		panic("unknown Payment method")
	}
	return constructor()
}
