package config

import (
	"cetm/qapi/config/cons"
	"cetm/qapi/x/db/mgo"
	"log"
	"os"
)

var db dbConfig

type dbConfig struct {
	DBHost    string
	DBName    string
	DBAccount string
}

func (db *dbConfig) check() bool {
	var _, err = mgo.NewDB(db.DBHost, db.DBName)
	if nil != err {
		log.Fatalln("Can't connect to db %s", db.DBHost)
		return false
	}
	os.Setenv(cons.ENV_OBJECT_DB, db.DBName)
	return true
}
