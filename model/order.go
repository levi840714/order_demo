package model

import (
	"time"
)

const (
	OrderStatusOk   = "1"
	OrderStatusFail = "0"
)

type Order struct {
	ID        int       `gorm:"column:id"`
	AccountId int       `gorm:"column:accountId"`
	GoodsId   int       `gorm:"column:goodsId"`
	Status    string    `gorm:"column:status"`
	OrderTime time.Time `gorm:"column:orderTime"`
}

func (Order) TableName() string {
	return "order"
}

func NewOrder(accountId int, goodsId int) (int, error) {
	insert := Order{AccountId: accountId, GoodsId: goodsId, Status: OrderStatusOk, OrderTime: time.Now()}
	if err := DB.Create(&insert).Error; err != nil {
		return 0, err
	}
	return insert.ID, nil
}

func DeleteOrder(id int) error {
	delete := Order{ID: id}
	if err := DB.Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}

type OrderInfo struct {
	ID        int       `gorm:"column:id" json:"id"`
	Goods     string    `gorm:"column:goods" json:"goods"`
	Amount    float64   `gorm:"column:amount" json:"amount"`
	Status    string    `gorm:"column:status" json:"status"`
	OrderTime time.Time `gorm:"column:orderTime" json:"orderTime"`
}

func GetTodayOrder(accountId int) (*[]OrderInfo, error) {
	var orders []OrderInfo
	today := time.Now().Format("2006-01-02")
	if err := DB.Table("order").Joins("JOIN goods ON goods.id = order.goodsId").
		Select("goods.goods, goods.amount, order.id, order.status, order.orderTime").
		Where("accountId = ? AND orderTime BETWEEN ? AND ?", accountId, today+" 00:00:00", today+" 23:59:59").Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}
