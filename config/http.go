package config

import "os"

type HttpConf struct {
	Addr string
}

func GetHttpConfig() *HttpConf {
	return &HttpConf{Addr: os.Getenv("HTTP_ADDR")}
}
