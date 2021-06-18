package model

import (
	"fmt"
	"log"
)

//图片
type Image struct {
	ID uint `gorm:"primary_key" json:"id"`
	//Title  string `json:"title"`
	URL string `json:"url"`
	//Width  uint   `json:"width"`
	//Height uint   `json:"height"`
}

func (Image) TableName() string {
	return "image"
}

/**
有匹配记录返回对应Id；否则新增记录。
*/
func (i *Image) AddImage() (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.AddImage: %w", err)
		}
	}()
	// 查找
	var count int64
	var image Image
	// 寻找URL一致的记录
	result := DB.Where(i).First(&image).Count(&count)
	// 没有一致的url
	if count == 0 {
		// 新增记录
		result = DB.Create(i)
		if result.Error != nil {
			log.Println(result.Error)
			return 0, result.Error
		}
	}
	// 已有该记录，返回对应ID
	return image.ID, err
}
