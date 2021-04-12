package controller

import (
	"github.com/gin-gonic/gin"
)

func GetBusinessById(ctx *gin.Context) {
	var data struct{
		ID string
	}
	ctx.ShouldBindJSON(&data)

}

func AddBusiness(ctx *gin.Context)  {
	var data struct{
		Name string
		Info string
	}
	ctx.ShouldBindJSON(&data)

}

func PostBusiness(ctx *gin.Context) {

}