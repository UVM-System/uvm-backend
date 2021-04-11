package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"uvm-backend/config"
	"uvm-backend/model"
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
	err = db.AutoMigrate(&model.Business{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Machine{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Model{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Order{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	once.Do(initDatabase)
	return db
}
