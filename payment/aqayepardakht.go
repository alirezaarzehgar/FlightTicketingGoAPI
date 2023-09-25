package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BaseMax/FlightTicketingGoAPI/config"
)

const PAYMENT_CALLBACK = "/payments/done"

type AqayePardakht struct {
	sandbox         bool
	pin             string
	createUrl       string
	verifyUrl       string
	callbackBaseURL string
}

func NewAqayePardakht(pin string, path string) *AqayePardakht {
	sandbox := false
	if pin == "sandbox" {
		sandbox = true
	}
	return &AqayePardakht{
		sandbox:         sandbox,
		pin:             pin,
		callbackBaseURL: config.GetPaymentCallback(),
		createUrl:       "https://panel.aqayepardakht.ir/api/v2/create",
		verifyUrl:       "https://panel.aqayepardakht.ir/api/v2/verify",
	}
}

func (gw AqayePardakht) Request(amount, transId uint) (string, error) {
	body, _ := json.Marshal(map[string]any{
		"pin":      gw.pin,
		"amount":   amount,
		"callback": fmt.Sprintf("%s/%d", gw.callbackBaseURL, transId),
	})
	resp, err := http.Post(gw.createUrl, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	var data map[string]any
	json.NewDecoder(resp.Body).Decode(&data)
	return data["transid"].(string), nil
}

func (gw AqayePardakht) CreateRequestUrl(authority string) string {
	prefix := ""
	if gw.sandbox {
		prefix = "sandbox/"
	}
	return fmt.Sprint("https://panel.aqayepardakht.ir/startpay/", prefix, authority)
}

func (gw AqayePardakht) Veify(amount uint, authority string) (bool, error) {
	return false, nil
}
