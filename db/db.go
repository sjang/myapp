package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"myapp/config"
)

var db *gorm.DB

func InitDB(cf *config.Config) {
	dataSourceName := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8mb4&parseTime=True&loc=Local", cf.DB.Username, cf.DB.Password, cf.DB.ServerHost, cf.DB.ServerPort, cf.DB.DBName)
	fmt.Println(dataSourceName)
	dbHandle, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	dbHandle.DB().SetMaxOpenConns(cf.DB.MaxOpenConns)
	dbHandle.DB().SetMaxIdleConns(cf.DB.MaxIdleConns)
	db = dbHandle
}

func DB() *gorm.DB {
	return db
}

func releaseDB() {
	db.Close()
}
