package router

import (
	"github.com/gin-gonic/gin"
	. "uvm-backend/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 用户
	r.POST("/user/wxLogin", WXLogin)
	r.GET("/user/getUserInfo", GetUserInfo)
	// 商家
	r.POST("/business/add", AddBusiness)
	r.GET("/business/getById", GetBusinessById)
	r.POST("/business/delete", DeleteBusiness)
	r.POST("/business/update", UpdateBusiness)
	// 商品
	r.GET("/product/productList", GetProductList)
	r.POST("/product/add", AddProduct)
	r.POST("/product/update", UpdateProduct)
	r.GET("/product/getInfoByEN", GetProductInfoByEN)
	// 图片下载
	r.GET("product/image/download", Download)

	return r
}
