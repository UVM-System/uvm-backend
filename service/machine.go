package service

import (
	"fmt"
	"log"
	"time"
	"uvm-backend/model"
)

func AddMachine(name, info string) (id uint, err error) {
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
