package payment

func GetPaymentWay(way string) Payment {
	switch way {
	case "CREDIT_CARD":
		return CreditCardPayment{}
	case "UPI":
		return UPIPayment{}
	case "NET_BANKING":
		return NetBanking{}
	default:
		panic("NO Sucha method exst as of now")

	}
}
