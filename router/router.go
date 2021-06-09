package router

import (
	"github.com/gin-gonic/gin"
	. "uvm-backend/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 用户登录
	r.POST("/auth/wxLogin", WXLogin)
	r.POST("/auth/getUserInfo", GetUserInfo)
	// 商家
	r.POST("/business/add", AddBusiness)
	r.GET("/business/getById", GetBusinessById)
	r.POST("/business/delete", DeleteBusiness)
	r.POST("/business/update", UpdateBusiness)
	// 商品
	r.POST("/product/productList", GetProductList)
	r.POST("/product/add", AddProduct)

	return r
}
