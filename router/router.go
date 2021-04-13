package router

import (
	"github.com/gin-gonic/gin"
	. "uvm-backend/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/business/add", AddBusiness)
	r.GET("/business/getById", GetBusinessById)
	return r
}
