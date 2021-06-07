package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"uvm-backend/service"
)

// 返回该商家的产品列表
func GetProductList(ctx *gin.Context) {
	var data struct {
		BusinessId uint
	}
	ctx.ShouldBindJSON(&data)
	productList, err := service.GetProductList(data.BusinessId)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message":      "读取商品列表成功",
		"product list": productList,
	})
}

// 增加产品
func AddProduct(ctx *gin.Context) {
	var data struct {
		BusinessId uint
		Name       string
		Info       string
		//Number     int
		//UpdateTime time.Time
		Price float64
		//Image 	   model.Image
	}
	ctx.ShouldBindJSON(&data)
	//id, err := service.AddProduct(data.BusinessId, data.Name, data.Info, data.Price, data.Image)
	id, err := service.AddProduct(data.BusinessId, data.Name, data.Info, data.Price)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "商品增加成功",
		"id":      id,
	})
}
