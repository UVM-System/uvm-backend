package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"strconv"
	"uvm-backend/config"
	"uvm-backend/database"
	"uvm-backend/model"
	"uvm-backend/router"
	"uvm-backend/service"
	"uvm-backend/utils"
)

/**
手动添加数据：1个商家，1个售货柜，100类商品图像及信息
！！重复执行会添加重复数据！！
*/
func initDB() {
	//// 1个商家
	log.Println("初始化商家……")
	businessId, err := service.AddBusiness("无人零售柜", "数据初始化测试")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("商家ID	", businessId)
	//// 1个售货柜，售货柜添加接口还没写

	//// 100类商品图像
	// 获取图像文件名list
	imgDir := "./upload/img/"
	imgList, err := utils.ReadDir(imgDir)
	if err != nil {
		log.Fatal(err)
	}
	// 图片id和path Map
	imgMap := make(map[string]uint)
	log.Println("初始化商品图像……")
	for _, fileName := range imgList {
		imgPath := imgDir + fileName.Name()
		imgId, err := service.AddImage(imgPath)
		if err != nil {
			log.Fatal(err)
		}
		englishName, _ := utils.GetFileNameAndSuffix(fileName.Name())
		// 英文名和图像id映射
		imgMap[englishName] = imgId
		log.Println(imgId, "	", imgPath)
	}
	//// 100类商品信息
	// 从excel文件的Sheet1中读取，列名：englishName, name, info, price
	log.Println("初始化商品信息……")
	file, err := excelize.OpenFile("./bev_name.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	// 读取Sheet1的所有行，每一行有两列：英文名和中文名
	rows := file.GetRows("Sheet1")
	for _, row := range rows {
		englishName := row[0]
		name := row[1]
		info := row[1]
		priceStr := "3"
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Fatal(err)
		}
		productId, err := service.AddProduct(businessId, name, englishName, info, price, imgMap[englishName])
		log.Println(productId, "	", businessId, "	", name, "	  ", englishName, "  	", info, "	  ", price, "	 ", imgMap[englishName])
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

func main() {
	//initDB() // 重复执行会添加重复数据！
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
