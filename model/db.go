package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"uvm-backend/config"
)

var once sync.Once
var db *gorm.DB

func initDatabase() {
	var err error
	conf := config.GetConfig()
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", conf.DBUsername, conf.DBPassword, conf.DBAddress, conf.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Business{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Machine{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Model{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Order{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	once.Do(initDatabase)
	return db
}
