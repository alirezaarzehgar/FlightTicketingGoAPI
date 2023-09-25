package payment

type AqayePardakht struct {
	sanbox    bool
	createUrl string
	verifyUrl string
}

func NewAqayePardakht(sandbox bool) *AqayePardakht {
	return &AqayePardakht{
		sanbox:    sandbox,
		createUrl: "https://panel.aqayepardakht.ir/api/v2/create",
		verifyUrl: "https://panel.aqayepardakht.ir/api/v2/verify",
	}
}

func (gw AqayePardakht) Request(amount uint) (string, error) {
	return "", nil
}

func (gw AqayePardakht) Veify(amount uint, authority string) (bool, error) {
	return false, nil
}
