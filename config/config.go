package config

import (
	"fmt"
	"os"
)

type allConfig struct {
	Survey SurveyHostConfig
	Host   hostConfig
	DB     dbConfig
	Email  EmailConfig
}

func (a *allConfig) check() {

	if !a.DB.check() {
		fmt.Println("Cant init db")
		os.Exit(1)
	}
	a.Host.check()
}
