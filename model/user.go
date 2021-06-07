package model

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	Id           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"size:50"`
	DealTimes    int       `json:"deal_times" gorm:"default:0;not null"`
	BusinessId   uint      `json:"business_id" gorm:"not null"`
	LastDealTime time.Time `json:"last_deal_time"`
	WXOpenId     string    `json:"wx_open_id" gorm:"not null;index:openid_idx"`
}

func (User) TableName() string {
	return "user"
}

/**
根据WXOpenId索引User对象
*/
func (u *User) GetUserByOpenId() (user User, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.GetUserByOpenId: %w", err)
		}
	}()
	result := DB.First(&user, u.WXOpenId)
	err = result.Error
	if err != nil {
		//log.Println(err)
		return User{}, err
	}
	return
}

/**
新增用户
*/
func (u *User) AddUser() (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("model.AddUser: %w", err)
		}
	}()
	result := DB.Create(u)
	if result.Error != nil {
		log.Println(result.Error)
		return 0, result.Error
	}
	return u.Id, err
}
