package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"uvm-backend/service"
)

/**
根据商品的英文名查找商品，并返回信息
*/
func GetProductInfoByEN(ctx *gin.Context) {
	englishName := ctx.Query("EnglishName")
	product, err := service.GetProductInfoByEN(englishName)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "根据英文名查找商品成功",
		"product": product,
	})
}

/**
返回某商家的所有产品列表
*/
func GetProductList(ctx *gin.Context) {
	businessId, err := strconv.Atoi(ctx.Query("BusinessId"))
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	_, _, _, productList, err := service.GetBusinessById(uint(businessId)) // 获取该商家的产品列表
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message":      "读取商品列表成功",
		"product_list": productList,
	})
}

/**
增加商品
*/
func AddProduct(ctx *gin.Context) {
	// 读取form data
	businessId, _ := strconv.Atoi(ctx.PostForm("BusinessId"))
	name := ctx.PostForm("Name")
	englishName := ctx.PostForm("EnglishName")
	info := ctx.PostForm("Info")
	price, _ := strconv.ParseFloat(ctx.PostForm("Price"), 64)
	log.Println("BusinessId: ", businessId, " Name: ", name, "EnglishName: ", englishName, " Info: ", info, " Price: ", price)
	// 读取文件
	file, header, err := ctx.Request.FormFile("upload")
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	filename := header.Filename
	log.Println(file, err, filename)

	// 把文件存入/upload/img/filename路径下
	filePath := "./upload/img/" + filename
	err = ctx.SaveUploadedFile(header, filePath)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	// 将图片信息存入数据库
	imgId, err := service.AddImage(filePath)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	// 将商品信息存入数据库
	productId, err := service.AddProduct(uint(businessId), name, englishName, info, price, imgId)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "商品增加成功",
		"id":      productId,
	})
}

/**
更新商品
*/
func UpdateProduct(ctx *gin.Context) {
	// 读取form data
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	businessId, _ := strconv.Atoi(ctx.PostForm("BusinessId"))
	name := ctx.PostForm("Name")
	englishName := ctx.PostForm("EnglishName")
	info := ctx.PostForm("Info")
	price, _ := strconv.ParseFloat(ctx.PostForm("Price"), 64)
	log.Println("id: ", id, " BusinessId: ", businessId, " Name: ", name, "EnglishName: ", englishName, " Info: ", info, " Price: ", price)
	// 读取文件
	file, header, err := ctx.Request.FormFile("upload")
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	filename := header.Filename
	log.Println(file, err, filename)

	// 把文件存入/upload/img/filename路径下
	filePath := "./upload/img/" + filename
	err = ctx.SaveUploadedFile(header, filePath)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	// 将图片信息存入数据库
	imgId, err := service.AddImage(filePath)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	// 更新数据库中的商品信息
	productId, err := service.UpdateProduct(uint(id), uint(businessId), name, englishName, info, price, imgId)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "商品更新成功",
		"id":      productId,
	})
}
