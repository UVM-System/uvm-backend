package service

import (
	"fmt"
	"log"
	"time"
	"uvm-backend/model"
)

func GetProductInfoByEN(englishName string) (p model.Product, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.GetProductInfoByEN: %w", err)
		}
	}()
	product := &model.Product{
		EnglishName: englishName,
	}
	p, err = product.GetProductByStructQuery()
	if err != nil {
		log.Println(err)
		return model.Product{}, err
	}
	return
}

/**
增加商品
*/
func AddProduct(businessId uint, name string, englishName string, info string, price float64, imageUrl string) (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.AddProduct: %w", err)
		}
	}()
	product := &model.Product{
		BusinessId:  businessId,
		Name:        name,
		EnglishName: englishName,
		Info:        info,
		Number:      0,
		UpdateTime:  time.Now(),
		Price:       price,
		ImageUrl:    imageUrl,
	}
	id, err = product.AddProduct()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return
}

/**
更新商品
*/
func UpdateProduct(id uint, businessId uint, name string, englishName string, info string, price float64, imageUrl string) (productId uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.UpdateProduct: %w", err)
		}
	}()
	product := &model.Product{
		Id:          id,
		BusinessId:  businessId,
		Name:        name,
		EnglishName: englishName,
		Info:        info,
		UpdateTime:  time.Now(),
		Price:       price,
		ImageUrl:    imageUrl,
	}
	productId, err = product.UpdateProduct()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return
}
