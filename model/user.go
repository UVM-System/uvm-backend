package model

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	Id           uint       `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name" gorm:"size:50"`
	DealTimes    int        `json:"deal_times" gorm:"default:0;not null"`
	BusinessId   uint       `json:"business_id" gorm:"not null";sql:"type:integer constraint fk_product_business REFERENCES business(id)"`
	Business     Business   `json:"business" gorm:"ForeignKey:BusinessId;AssociationForeignKey:ID"`
	LastDealTime *time.Time `json:"last_deal_time"` // *time.Time允许空值
	//SessionId     string    `json:"session_id" gorm:"not null;index:openid_idx"`
	WXOpenId  string `json:"session_id" gorm:"not null;index:openid_idx"`
	Nickname  string `json:"nickName" gorm:"size:50"`
	AvatarUrl string `json:"avatarUrl"`
}

func (User) TableName() string {
	return "user"
}

//// 第三方自定义Session
//type CustomSession struct {
//	OpenId 		string		`json:"openid"`
//	SessionKey  string 		`json:"session_key"`
//}

/**
更新User记录
*/
func (u *User) UpdateUser() (user User, err error) {
	defer func() {
		if err != nil && err != gorm.ErrRecordNotFound {
			// 没有记录时不修改，否则service不好判断
			err = fmt.Errorf("model.UpdateUser: %w", err)
		}
	}()
	result := DB.Model(u).Updates(*u)
	err = result.Error
	if err != nil {
		return User{}, err
	}
	return *u, nil
}

/**
根据WXOpenId或者Id（候选键）索引User对象
*/
func (u *User) GetUserByID() (user User, err error) {
	defer func() {
		if err != nil && err != gorm.ErrRecordNotFound {
			// 没有记录时不修改，否则service不好判断
			err = fmt.Errorf("model.GetUserByOpenId: %w", err)
		}
	}()
	result := DB.Where(u).First(&user)
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
