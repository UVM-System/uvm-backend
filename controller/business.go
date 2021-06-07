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
		"name":          name,
		"info":          info,
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
		"message": "增加成功",
		"id":      id,
		"name":    data.Name,
		"info":    data.Info,
	})
}

func DeleteBusiness(ctx *gin.Context) {
	var data struct {
		ID uint
	}
	ctx.ShouldBindJSON(&data)
	err := service.DeleteBusiness(data.ID)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "删除成功",
	})
}

func UpdateBusiness(ctx *gin.Context) {
	var data struct {
		ID   uint
		Name string
		Info string
	}
	ctx.ShouldBindJSON(&data)
	name, info, t, err := service.UpdateBusiness(data.ID, data.Name, data.Info)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message":       "更新成功",
		"id":            data.ID,
		"name":          name,
		"info":          info,
		"register time": t,
	})
}
