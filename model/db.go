package model

import (
	"uvm-backend/database"
)

/**
MySQL
*/
var DB = database.GetDB()

func init() {
	err := DB.AutoMigrate(&Business{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Machine{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Model{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Order{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Goods{})
	if err != nil {
		panic(err)
	}
}
