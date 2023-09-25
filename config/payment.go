package config

import "os"

func GetPaymentCallback() string {
	return os.Getenv("PAYMENT_CALLBACK_URL")
}
