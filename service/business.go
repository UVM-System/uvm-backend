package service

import (
	"log"
	"time"
	"uvm-backend/model"
)

// 增加商家
// @param name string "商家名称"
// @param info string "商家信息"
// @return id uint "商家id"
func AddBusiness(name, info string) (id uint, err error) {
	business := &model.Business{
		Name:         name,
		Info:         info,
		RegisterTime: time.Now(),
	}
	id, err = business.AddBusiness()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return
}
