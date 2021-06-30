package service

import (
	"fmt"
	"log"
	"uvm-backend/model"
)

type GoodsInfo struct {
	//售货柜对应商品的商品详细信息
	ProductId uint    `json:"product_id"` // 产品编号
	Name      string  `json:"name"`
	Info      string  `json:"info"`
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	ImageUrl  string  `json:"image_url"`
}

func AddMachine(businessId uint, location string, modelId uint) (id uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.AddMachine: %w", err)
		}
	}()

	machine := &model.Machine{
		BusinessId: businessId,
		Location:   location,
		ModelId:    modelId,
	}
	id, err = machine.AddMachine()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return
}

/**
根据售货柜Id，返回售货柜商品完整信息
*/
func GetMachineGoodsById(id uint) (businessId uint, location string, goodsInfoList []GoodsInfo, modelId uint, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("service.GetMachineById: %w", err)
		}
	}()
	machine := &model.Machine{
		Id: id,
	}
	m, err := machine.GetMachineGoodsById()
	if err != nil {
		log.Println(err)
		return 0, "", nil, 0, err
	}
	// goods中只保存了售货柜中商品数量这一信息，故查询对应的product，将完整信息返回
	for _, goods := range m.Goods {
		product := &model.Product{
			Id: goods.ProductId,
		}
		p, err := product.GetProductByStructQuery()
		if err != nil {
			log.Println(err)
			return 0, "", nil, 0, err
		}
		goodsInfo := GoodsInfo{goods.ProductId, p.Name, p.Info, goods.Number, p.Price, p.ImageUrl}
		goodsInfoList = append(goodsInfoList, goodsInfo)
	}
	// 根据goods
	return m.BusinessId, m.Location, goodsInfoList, m.ModelId, nil
}
