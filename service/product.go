package service

import (
	"fmt"
	"log"
	"time"
	"uvm-backend/model"
)

// 返回商家的商品列表
func GetProductList(businessId uint) (productList []model.Product, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.GetProductList: %w", err)
		}
	}()
	productList, err = model.GetProductList(businessId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

// 增加商品
//func AddProduct(businessId uint, name string, info string, price float64, image model.Image) (id uint, err error){
func AddProduct(businessId uint, name string, info string, price float64) (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.AddProduct: %w", err)
		}
	}()
	product := &model.Product{
		BusinessId: businessId,
		Name:       name,
		Info:       info,
		Number:     0,
		UpdateTime: time.Now(),
		Price:      price,
		//Image: image,
	}
	id, err = product.AddProduct()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return
}
