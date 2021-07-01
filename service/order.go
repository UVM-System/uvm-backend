package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"uvm-backend/model"
)

type OrderInfo struct {
	// 解析后的订单信息
	UserId      uint      `json:"user_id"`
	BusinessId  uint      `json:"business_id"`
	MachineId   uint      `json:"machine_id"`
	OrderNumber string    `json:"order_number"`
	Status      bool      `json:"status"`
	Price       float64   `json:"price"`
	Number      int       `json:"number"`
	Items       []Item    `json:"items"`
	CreatedTime time.Time `json:"created_time"`
}
type Item struct {
	// 订单内容解析
	ProductId uint    `json:"product_id"`
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	Name      string  `json:"name"`
}

/**
将Orders解析为OrderInfo
*/
func orderListParse(orders []model.Order) (orderInfos []OrderInfo, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("srevice.orderListParse: %w", err)
		}
	}()
	for _, order := range orders {
		orderInfo, err := orderParse(order)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		orderInfos = append(orderInfos, orderInfo)
	}
	return orderInfos, nil
}
func orderParse(order model.Order) (orderInfo OrderInfo, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("srevice.orderParse: %w", err)
		}
	}()
	orderContent := order.OrderContent
	strList := strings.Split(orderContent, "|") // 切割字符串, productId|number|price|name
	// 解析订单内容
	totalNumber := 0
	var items []Item
	for i := 0; i < len(strList); i += 4 {
		productId, err := strconv.Atoi(strList[i])
		if err != nil {
			log.Println(err)
			return OrderInfo{}, err
		}
		number, err := strconv.Atoi(strList[i+1])
		if err != nil {
			log.Println(err)
			return OrderInfo{}, err
		}
		price, err := strconv.ParseFloat(strList[i+2], 64)
		if err != nil {
			log.Println(err)
			return OrderInfo{}, err
		}
		name := strList[i+3]
		item := Item{uint(productId), number, price, name}
		items = append(items, item)
		totalNumber += number
	}
	orderInfo = OrderInfo{order.UserId, order.BusinessId, order.MachineId, order.OrderNumber, order.Status, order.Price, totalNumber, items, order.CreatedTime}
	return
}

/**
根据订单号查询订单
*/
func GetOrderByOrderNumber(orderNumber string) (orderInfo OrderInfo, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("srevice.GetOrderByOrderNumber: %w", err)
		}
	}()
	order := model.Order{
		OrderNumber: orderNumber,
	}
	o, err := order.GetOrderByStructQuery()
	orderInfo, err = orderParse(o)
	if err != nil {
		log.Println(err)
		return OrderInfo{}, err
	}
	return
}

/**
查询某商家某日期及某状态的订单列表
*/
func GetOrderByDateNStatus(businessId uint, date string, status bool) (orderInfos []OrderInfo, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("srevice.GetOrderByDateNStatus: %w", err)
		}
	}()
	// 求时间区间
	timeLayout := "2006-01-02"                                           // 时间格式
	startTime, err := time.ParseInLocation(timeLayout, date, time.Local) // 当天零点
	if err != nil {
		log.Println(err)
		return nil, err
	}
	endTime := startTime.AddDate(0, 0, 1) // 第二天零点
	log.Println("startTime: ", startTime, " endTime: ", endTime)
	order := &model.Order{
		BusinessId: businessId,
		Status:     status,
	}
	orders, err := order.GetOrderListWithinSpan(businessId, status, startTime, endTime)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// 将orders解析为orderInfo
	orderInfos, err = orderListParse(orders)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}
func GetOrdersByUserId(userId uint) (orderInfos []OrderInfo, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("srevice.GetOrdersByUserId: %w", err)
		}
	}()
	o := &model.Order{
		UserId: userId,
	}
	orders, err := o.GetOrderListByStructQuery()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// 将orders解析为orderInfo
	orderInfos, err = orderListParse(orders)
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
