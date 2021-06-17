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

func (i *Image) AddImage() (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.AddImage: %w", err)
		}
	}()
	result := DB.Create(i)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}
	return i.ID, err
}
