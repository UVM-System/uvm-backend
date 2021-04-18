package service

import (
	"fmt"
	"log"
	"time"
	"uvm-backend/model"
)

// 增加商家
// @param name string "商家名称"
// @param info string "商家信息"
// @return id uint "商家id"
func AddBusiness(name, info string) (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.AddBusiness: %w", err)
		}
	}()
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

func GetBusinessById(id uint) (name, info string, t time.Time, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.GetBusinessById: %w", err)
		}
	}()
	business := &model.Business{
		ID: id,
	}
	b, err := business.GetBusinessById()
	if err != nil {
		log.Println(err)
		return "", "", time.Now(), err
	}
	return b.Name, b.Info, b.RegisterTime, nil
}
