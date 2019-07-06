package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"myapp/config"
	"myapp/db"
	"myapp/logging"
	"os"
)

// static build
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a main.go

var log *logrus.Logger
var cf *config.Config
var dbHandle *gorm.DB

func init() {
	// checking for init

	// config(json)
	initConfig()

	// init logrus for logging
	logging.InitLog(cf)
	log = logging.Log()

	// DB
	initDB(cf)
}

func initConfig() {
	configFlag := flag.Bool("f", false, "-f myapp.conf")
	flag.Parse()
	var configFilepath string
	if *configFlag == true {
		configFilepath = flag.Args()[0]
	} else {
		fmt.Println("Fail to load config...")
		os.Exit(1)
	}

	err := config.LoadConfig(configFilepath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("config LOADED!(%s)", configFilepath)

	cf = config.Conf()
	fmt.Println(cf)
}

func initDB(cf *config.Config) {
	db.InitDB(cf)
	dbHandle = db.DB()
}

// main
func main() {
	log.WithFields(logrus.Fields{
		"testKey": "testValue",
	}).Info("start the app")

	fmt.Println("Hello myapp!")
}
