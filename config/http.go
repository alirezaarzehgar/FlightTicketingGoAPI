package config

import "os"

type HttpConf struct {
	Addr string
}

func GetHttpConfig() *HttpConf {
	return &HttpConf{Addr: os.Getenv("HTTP_ADDR")}
}

func GetJwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func GetAirlineConfPath() string {
	return os.Getenv("AIRLINE_CONF_PATH")
}
