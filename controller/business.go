package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"uvm-backend/service"
)

func GetBusinessById(ctx *gin.Context) {
	var data struct {
		ID string
	}
	ctx.ShouldBindJSON(&data)

}

func AddBusiness(ctx *gin.Context) {
	var data struct {
		Name string
		Info string
	}
	ctx.ShouldBindJSON(&data)
	log.Println("name: ", data.Name, "\t info: ", data.Info)
	id, err := service.AddBusiness(data.Name, data.Info)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "增加失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "增加成功",
		"id":      id,
	})
}

func PostBusiness(ctx *gin.Context) {

}
