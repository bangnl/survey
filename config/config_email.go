package config

import (
	"fmt"
)

//EmailConfig :
type EmailConfig struct {
	Port     int
	Host     string
	Email    string
	Pwd      string
	Receives []string
}

func (ec *EmailConfig) check() {
	if ec.Port == 0 {
		fmt.Println("WARNING: Email config missing mail server port")
	}
	if ec.Email == "" || ec.Pwd == "" || ec.Host == "" {
		fmt.Println("WARNING: Email config missing mail config")
	}
	if len(ec.Receives) == 0 {
		fmt.Println("WARNING: Email config have no one get mail")
	}
}
