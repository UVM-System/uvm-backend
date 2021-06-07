package main

import (
	"fmt"
	"log"
	"uvm-backend/config"
	"uvm-backend/database"
	"uvm-backend/model"
	"uvm-backend/router"
)

func main() {
	conf := config.GetConfig()
	if err := router.Router().Run(":" + conf.PORT); err != nil {
		log.Fatal(err)
	}

}

func ope() (err error) {
	db := database.GetDB()
	trans := db.Begin()
	defer func() {
		if err != nil {
			trans.Callback()
			err = fmt.Errorf("ope: %w", err)
		} else {
			trans.Commit()
		}
	}()
	rows, err := trans.Raw("select * from business").Rows()
	if err != nil {
		return err
	}
	var bus model.Business
	for rows.Next() {
		err := db.ScanRows(rows, &bus)
		if err != nil {
			return err
		}
		log.Println(bus.Name)
	}
	//err = trans.Raw("select * from businesses where name = ?", "丑八怪").First(&bus).Error
	//if err != nil {
	//	return err
	//}
	log.Println("first", bus.Name)
	return nil
}
