package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path"
)

/**
服务器端存储图片下载接口
*/
func Download(ctx *gin.Context) {
	// *匹配不到./……手动加上（
	filePath := "./" + ctx.Param("path")
	fileName := path.Base(filePath)
	log.Println(filePath)
	log.Println(fileName)
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.File(filePath)
	return
}
