package config

import (
	"fmt"
	"log"
	"os"
)

type appConfig struct {
	serverAddress string
	serverPort    string
	dbDialect     string
	dbUser        string
	dbPasswd      string
	dbPort        string
	dbName        string
}

var appConf *appConfig

func envSanityCheck(env string) string {
	if os.Getenv(env) == "" {
		log.Fatalln("Unable to find environment variable: " + env)
	}
	return os.Getenv(env)
}

func InitAppConfig() {
	appConf = &appConfig{
		serverAddress: envSanityCheck("SERVER_ADDRESS"),
		serverPort:    envSanityCheck("SERVER_PORT"),
		dbDialect:     envSanityCheck("DB_DIALECT"),
		dbUser:        envSanityCheck("DB_USER"),
		dbPasswd:      envSanityCheck("DB_PASSWD"),
		dbPort:        envSanityCheck("DB_PORT"),
		dbName:        envSanityCheck("DB_NAME"),
	}

}

func GetServerAddr() string {
	return appConf.serverAddress + ":" + appConf.serverPort
}

func GetDatabaseSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s",
		appConf.dbUser,
		appConf.dbPasswd,
		appConf.dbPort,
		appConf.dbName,
	)
}

func GetDbDialect() string {
	return appConf.dbDialect
}
