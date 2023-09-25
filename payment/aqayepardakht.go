package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AqayePardakht struct {
	sandbox     bool
	pin         string
	createUrl   string
	verifyUrl   string
	callbackURL string
}

func NewAqayePardakht(pin string, path string) *AqayePardakht {
	sandbox := false
	if pin == "sandbox" {
		sandbox = true
	}
	return &AqayePardakht{
		sandbox:     sandbox,
		pin:         pin,
		callbackURL: path + PAYMENT_CALLBACK,
		createUrl:   "https://panel.aqayepardakht.ir/api/v2/create",
		verifyUrl:   "https://panel.aqayepardakht.ir/api/v2/verify",
	}
}

func (gw AqayePardakht) Request(amount uint) (string, error) {
	body, _ := json.Marshal(map[string]any{
		"pin":      gw.pin,
		"amount":   amount,
		"callback": gw.callbackURL,
	})
	resp, err := http.Post(gw.createUrl, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	var data map[string]any
	json.NewDecoder(resp.Body).Decode(&data)
	prefix := ""
	if gw.sandbox {
		prefix = "sandbox/"
	}
	url := fmt.Sprint("https://panel.aqayepardakht.ir/startpay/", prefix, data["transid"])
	return url, nil
}

func (gw AqayePardakht) Veify(amount uint, authority string) (bool, error) {
	return false, nil
}
