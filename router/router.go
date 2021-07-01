package router

import (
	"github.com/gin-gonic/gin"
	. "uvm-backend/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	//// 用户
	//登录
	r.POST("/user/wxLogin", WXLogin)
	r.GET("/user/getUserInfo", GetUserInfo)
	// 订单
	r.POST("/user/order/add", AddOrder)
	r.GET("/user/order/orderList", GetOrdersByUserId)
	// 商品
	r.GET("/user/machine/goodsList", GetGoodsByMachineId)

	//// 商家
	r.POST("/business/add", AddBusiness)
	r.POST("/business/delete", DeleteBusiness)
	r.POST("/business/update", UpdateBusiness)
	// 商品
	r.GET("/business/product/productList", GetProductsByBusinessId)
	r.POST("/business/product/add", AddProduct)
	r.POST("/business/product/update", UpdateProduct)
	r.GET("/business/product/getInfoByEN", GetProductInfoByEN)
	// 订单
	r.GET("/business/order/orderList", GetOrdersByDateNStatus)
	r.GET("/business/order/detail", GetOrderByOrderNumber)
	// 售货柜
	r.POST("/business/machine/add", AddMachine)

	//// 商品图片下载
	r.GET("/product/image/download", Download)

	return r
}
