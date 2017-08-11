package config

import (
	"fmt"

	"os"

	"flag"

	"github.com/BurntSushi/toml"
)

const (
	currentVersion = "1.0.0-beta.01"
)

var configFile = "./static/survey.toml"
var all allConfig

//Init : init config
func init() {
	conf := flag.String("conf", configFile, "Config file")
	version := flag.Bool("version", false, "Current version")
	flag.Parse()

	if *version {
		fmt.Print(currentVersion)
		os.Exit(0)
	}

	_, err := toml.DecodeFile(*conf, &all)
	if err != nil {
		fmt.Println("Cant init config with file: ", *conf)
		os.Exit(1)
	}
	all.check()
	host = all.Host
	db = all.DB
}

//GetSurveyHost :
func GetSurveyHost() string {
	return all.Survey.Host
}

func GetEmailConfig() EmailConfig {
	return all.Email
}

func GetHostPort() string {
	return host.Port
}
