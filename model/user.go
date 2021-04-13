package model

import "time"

type User struct {
	Id           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"size:50"`
	DealTimes    int       `json:"deal_times" gorm:"default:0;not null"`
	BusinessId   uint      `json:"business_id" gorm:"not null"`
	LastDealTime time.Time `json:"last_deal_time"`
}

func (User) TableName() string {
	return "user"
}
