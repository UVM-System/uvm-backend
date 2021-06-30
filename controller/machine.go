package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"uvm-backend/service"
)

/**
增加售货柜
*/
func AddMachine(ctx *gin.Context) {
	var data struct {
		BusinessId uint
		Location   string
		ModelId    uint
	}
	ctx.ShouldBindJSON(&data)
	id, err := service.AddMachine(data.BusinessId, data.Location, data.ModelId)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "售货柜增加成功",
		"id":      id,
	})
}

/**
根据售货柜Id，获取商品列表
*/
func GetGoodsByMachineId(ctx *gin.Context) {
	machineId, err := strconv.Atoi(ctx.Query("MachineId"))
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	_, _, goodsInfoList, _, err := service.GetMachineGoodsById(uint(machineId))
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message":    "读取售货柜商品列表成功",
		"goods_list": goodsInfoList,
	})

}
