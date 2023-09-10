package config

import (
	"os"
	"strconv"
	"strings"
)

type DbConf struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     uint64
	TimeZone string
	Debug    bool
}

func GetDbConfig() (*DbConf, error) {
	port, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 64)
	if err != nil {
		return nil, err
	}

	debug := false
	if strings.ToLower(os.Getenv("DB_DEBUG")) == "true" {
		debug = true
	}

	conf := DbConf{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Port:     port,
		TimeZone: os.Getenv("DB_TZ"),
		Debug:    debug,
	}

	return &conf, nil
}
