package service

import (
	"fmt"
	"log"
	"uvm-backend/model"
)

/**
增加图片；
当存在数据库中已有图片url与待新增图片一致时，返回该图片Id；
反之新增图片记录。
*/
func AddImage(filePath string) (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.AddImage: %w", err)
		}
	}()
	image := &model.Image{
		URL: filePath,
	}
	id, err = image.AddImage()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return
}
