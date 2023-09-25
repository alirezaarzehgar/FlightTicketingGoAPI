package payment

const PAYMENT_CALLBACK = "/payments/done"

type PaymentGateway interface {
	Request(amount uint) (string, error)
	Veify(amount uint, authority string) (bool, error)
}

func CreateRequest(pgw PaymentGateway, amount uint) (string, error) {
	return pgw.Request(amount)
}

func Verify(pgw PaymentGateway, amount uint, authority string) (bool, error) {
	return pgw.Veify(amount, authority)
}
