package config

import "os"

type AdminConf struct {
	Email    string
	Password string
}

func GetAdminConf() *AdminConf {
	return &AdminConf{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
	}
}
