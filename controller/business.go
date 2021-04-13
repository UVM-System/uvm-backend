package controller

import (
	"github.com/gin-gonic/gin"
	"log"
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

}

func PostBusiness(ctx *gin.Context) {

}
