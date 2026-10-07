package config

import (
	"fmt"
	"os"
)

type DatabaseCredentials struct {
	Driver   string
	Password string
	User     string
	Host     string
	Uri      string
}

func NewDatabaseCredentials() *DatabaseCredentials {
	driveType := os.Getenv("DB_DRIVER")
	psw := os.Getenv("DB_PSW")
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	uri := fmt.Sprintf(`%s%s:%s@%s`, os.Getenv("DB_URI"), user, psw, host)

	return &DatabaseCredentials{
		Driver:   driveType,
		Password: psw,
		User:     user,
		Host:     host,
		Uri:      uri,
	}
}
