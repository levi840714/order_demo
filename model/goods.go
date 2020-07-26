package model

import (
	"time"
)

type Goods struct {
	ID       int       `gorm:"column:id"`
	Good     string    `gorm:"column:good`
	Amount   float64   `gorm:"column:amount"`
	Status   string    `gorm:"column:status"`
	CreateAt time.Time `gorm:"column:createAt"`
}

func (Goods) TableName() string {
	return "goods"
}

func GetGoods() {

}

func AddGoods(good string, amount float64) (int, error) {
	insert := Goods{Good: good, Amount: amount, Status: "1", CreateAt: time.Now()}
	if err := DB.Create(&insert).Error; err != nil {
		return 0, err
	}
	return insert.ID, nil
}

func UpdateGoods() {

}

func DeleteGoods(id int, status string) error {
	delete := Goods{ID: id}
	if err := DB.Model(&delete).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
