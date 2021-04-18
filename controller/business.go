package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"uvm-backend/service"
)

func GetBusinessById(ctx *gin.Context) {
	var data struct {
		ID uint
	}
	ctx.ShouldBindJSON(&data)
	name, info, t, err := service.GetBusinessById(data.ID)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"name": name,
		"info": info,
		"register time": t,
	})
}

func AddBusiness(ctx *gin.Context) {
	var data struct {
		Name string
		Info string
	}
	ctx.ShouldBindJSON(&data)
	id, err := service.AddBusiness(data.Name, data.Info)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"code":    1,
		"message": "增加成功",
		"id":      id,
	})
}
