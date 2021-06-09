package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"uvm-backend/service"
)

/**
用户保持登录态；根据用户ID获取avatarUrl和nickName
*/
func GetUserInfo(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.PostForm("userId"))
	log.Println(userId)
	avatarUrl, nickName, err := service.GetUserInfo(uint(userId))
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	SuccessResponse(ctx, gin.H{
		"message":   "读取用户基本信息成功",
		"avatarUrl": avatarUrl,
		"nickName":  nickName,
	})
}

/**
用户授权登录：保存用户openID、avatarUrl和nickName
*/
func WXLogin(ctx *gin.Context) {
	// 报文格式：x-www-form-urlencoded
	code := ctx.PostForm("code")
	avatarUrl := ctx.PostForm("avatarUrl")
	nickName := ctx.PostForm("nickName")
	log.Println(code)
	// 根据code获取敏感信息：openID和session_key
	wxLoginResp, err := service.GetWXSession(code)
	if err != nil {
		log.Println(err)
		ErrorResponse(ctx, err)
		return
	}
	// 根据openID创建/更新用户
	id, err := service.UserLogin(wxLoginResp.OpenID, avatarUrl, nickName)
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
