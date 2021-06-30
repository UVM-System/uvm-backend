package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"uvm-backend/service"
)

/**
通过UserId查询用户订单列表
*/
func GetOrdersByUserId(ctx *gin.Context) {
	UserId, err := strconv.Atoi(ctx.Query("UserId"))
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	Orders, err := service.GetOrdersByUserId(uint(UserId))
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "获取用户订单列表成功",
		"orders":  Orders,
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
