package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"uvm-backend/service"
)

/**
根据订单号查询订单
*/
func GetOrderByOrderNumber(ctx *gin.Context) {
	OrderNumber := ctx.Query("OrderNumber")
	order, err := service.GetOrderByOrderNumber(OrderNumber)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "获取订单成功",
		"order":   order,
	})
}

/**
通过UserId查询用户订单列表
*/
func GetOrdersByUserId(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("UserId"))
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	orders, err := service.GetOrdersByUserId(uint(userId)) // 按id排序
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "获取用户订单列表成功",
		"orders":  orders,
	})
}

/**
查询某商家具体日期和状态的交易订单列表
*/
func GetOrdersByDateNStatus(ctx *gin.Context) {
	businessId, err := strconv.Atoi(ctx.Query("BusinessId"))
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	date := ctx.Query("Date") // 格式形如2021-06-30
	// 订单状态
	statusInt, err := strconv.Atoi(ctx.Query("Status"))
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	// 转换为bool
	var status bool
	if statusInt != 0 {
		status = true
	} else {
		status = false
	}
	orders, err := service.GetOrderByDateNStatus(uint(businessId), date, status)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "获取指定日期及状态的商家订单列表成功",
		"orders":  orders,
	})

}

/**
添加订单
*/
func AddOrder(ctx *gin.Context) {
	var data struct {
		UserId       uint
		BusinessId   uint
		MachineId    uint
		Status       bool
		OrderNumber  string
		OrderContent string
		TotalPrice   float64
	}
	ctx.ShouldBindJSON(&data)
	log.Println(data)
	id, err := service.AddOrder(data.OrderNumber, data.OrderContent, data.TotalPrice, data.UserId, data.MachineId, data.BusinessId, data.Status)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "订单添加成功",
		"id":      id,
	})

}
