package model

import (
	"time"
)

const (
	GoodsStatusFail = "0"
	GoodsStatusOK   = "1"
)

type Goods struct {
	ID       int       `gorm:"column:id" json:"id"`
	Goods    string    `gorm:"column:goods" json:"goods"`
	Amount   float64   `gorm:"column:amount" json:"amount"`
	Status   string    `gorm:"column:status" json:"status"`
	CreateAt time.Time `gorm:"column:createAt" json:"createAt"`
}

func (Goods) TableName() string {
	return "goods"
}

func GetGoods(status string) (*[]Goods, error) {
	var goods []Goods
	if err := DB.Where("status = ?", status).Find(&goods).Error; err != nil {
		return nil, err
	}
	return &goods, nil

}

func AddGoods(goods string, amount float64) (int, error) {
	insert := Goods{Goods: goods, Amount: amount, Status: GoodsStatusOK, CreateAt: time.Now()}
	if err := DB.Create(&insert).Error; err != nil {
		return 0, err
	}
	return insert.ID, nil
}

func UpdateGoods(id int, goods string, amount float64) error {
	var goodsData Goods
	if err := DB.Where("id = ?", id).Find(&goodsData).Error; err != nil {
		return err
	}
	goodsData.Goods = goods
	goodsData.Amount = amount
	if err := DB.Save(&goodsData).Error; err != nil {
		return err
	}
	return nil

}

func DeleteGoods(id int) error {
	delete := Goods{ID: id}
	if err := DB.Model(&delete).Update("status", GoodsStatusFail).Error; err != nil {
		return err
	}
	return nil
}
