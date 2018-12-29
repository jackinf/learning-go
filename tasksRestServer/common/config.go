package common

import (
	"encoding/json"
	"log"
	"os"
)

var AppConfig configuration

type configuration struct {
	Server, MogoDbHost, DBUser, DBPwd, Database string
}

func initConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("/Users/jrumjantsev/dev/tmp/task-manager-config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig: %s\n", err)
	}

	AppConfig = configuration{}
	err = json.NewDecoder(file).Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
