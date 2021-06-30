package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"uvm-backend/model"
)

func GetOrdersByUserId(userId uint) (orders []model.Order, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("srevice.GetOrdersByUserId: %w", err)
		}
	}()
	o := &model.Order{
		UserId: userId,
	}
	orders, err = o.GetOrderListByStructQuery()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}

/**
添加新的订单，分为两部分：
1. 保存订单数据；
2. 根据订单调整Goods和Product对应记录的Number。
*/
func AddOrder(orderNumber string, orderContent string, price float64, userId uint, machineId uint, businessId uint, status bool) (Id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.AddOrder: %w", err)
		}
	}()
	// 保存订单
	order := &model.Order{
		UserId:       userId,
		BusinessId:   businessId,
		MachineId:    machineId,
		OrderNumber:  orderNumber,
		OrderContent: orderContent,
		Price:        price,
		Status:       status,
		CreatedTime:  time.Now(),
	}
	Id, err = order.AddOrder()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// 根据订单调整Goods、Product的Number
	pnStringlist := strings.Split(orderContent, "|") // 切割字符串, productId|number|price|name
	log.Println("pnStringlist: ", pnStringlist)
	for i := 0; i < len(pnStringlist); i += 4 {
		// 取出productId和对应的number
		productId, err := strconv.Atoi(pnStringlist[i])
		if err != nil {
			log.Println(err)
			return 0, err
		}
		number, err := strconv.Atoi(pnStringlist[i+1])
		if err != nil {
			log.Println(err)
			return 0, err
		}
		// 根据ProductId和MachineId找到对应Goods，调整Number
		_, err = UpdateGoods(uint(productId), machineId, number, true)
		if err != nil {
			log.Println(err)
			return 0, err
		}
		// 调整Product的Number
		_, err = UpdateProductNumber(uint(productId), number, true)
		if err != nil {
			log.Println(err)
			return 0, err
		}
	}
	return
}
