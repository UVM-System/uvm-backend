package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"uvm-backend/service"
)

/**
得到微信用户openID，进而进行用户登录
*/
func WXLogin(ctx *gin.Context) {
	// 报文格式：x-www-form-urlencoded
	code := ctx.PostForm("code")
	log.Println(code)
	// 根据code获取openID和session_key
	wxLoginResp, err := service.WXLogin(code)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	// 根据openID查找/创建用户
	id, err := service.UserLogin(wxLoginResp.OpenID)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message": "登录成功",
		"id":      id,
	})
}
