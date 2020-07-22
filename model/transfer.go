package model

import (
	"order_demo/lib/logger"
	"time"
)

type Transfer struct {
	ID           int       `gorm:"column:id"`
	AccountId    int       `gorm:"column:accountId"`
	Amount       float64   `gorm:"column:amount"`
	TransferTime time.Time `gorm:"column:transferTime"`
}

func (Transfer) TableName() string {
	return "transfer"
}

func AddTransfer(accountId int, amount float64) bool {
	insert := Transfer{AccountId: accountId, Amount: amount, TransferTime: time.Now()}
	if err := DB.Create(&insert).Error; err != nil {
		logger.Error.Println(err.Error())
		return false
	}
	return true
}
