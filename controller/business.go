package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"uvm-backend/service"
)

func GetBusiness(ctx *gin.Context) {
	var data struct{
		ID string
	}
	ctx.ShouldBindJSON(&data)
	var err error
	if errors.Is(err, service.ErrorNameInvalid){

	}
}

func PostBusiness(ctx *gin.Context) {

}