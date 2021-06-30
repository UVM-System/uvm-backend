package service

import (
	"fmt"
	"log"
	"uvm-backend/model"
)

/**
新增Goods；需要更新对应Product的number
*/
func AddGoods(machineId uint, productId uint, number int) (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.AddGoods: %w", err)
		}
	}()
	// 添加Goods对象
	goods := &model.Goods{
		MachineId: machineId,
		ProductId: productId,
		Number:    number,
	}
	id, err = goods.AddGoods()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// 增加Product的Number
	_, err = UpdateProductNumber(uint(productId), number, false)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return
}

/**
根据ProductId和MachineId查询Goods，更新Goods
change为库存改变量
isConsume为1时表示消费该商品，change为消耗量；为0时表示上货，change为上货量
*/
func UpdateGoods(productId uint, machineId uint, change int, isConsume bool) (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.GetGoodsByProductNMachineId: %w", err)
		}
	}()
	// 根据productId和machineId查找goods
	goods := &model.Goods{
		ProductId: productId,
		MachineId: machineId,
	}
	g, err := goods.GetGoodsByStructQuery()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// 更新goods的number
	var number int
	if isConsume {
		number = g.Number - change
	} else {
		number = g.Number + change
	}
	updatedGoods := &model.Goods{
		Id:     g.Id,
		Number: number,
	}
	updatedG, err := updatedGoods.UpdateGoods()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return updatedG.Id, nil
}
