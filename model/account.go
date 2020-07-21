package model

import (
	"time"
)

type Account struct {
	ID       int       `gorm:"column:id"`
	Account  string    `gorm:"column:account"`
	Password string    `gorm:"column:password"`
	Status   string    `gorm:"column:status"`
	Balance  float64   `gorm:"column:balance"`
	CreateAt time.Time `gorm:"column:createAt"`
}

func (Account) TableName() string {
	return "account"
}

func RegisterAccount(account string, password string) (code int, err error) {
	createData := Account{Account: account, Password: password, Status: "1", Balance: 0.0, CreateAt: time.Now()}
	if err := DB.Create(&createData).Error; err != nil {
		return 1, err
	}
	return 0, nil

}

func GetAccount(accountId int) (*Account, error) {
	var account Account
	if err := DB.Where("id = ?", accountId).Find(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func CheckLogin(account string, password string) bool {
	var login Account
	if err := DB.Where("account = ? AND password = ?", account, password).Find(&login).Error; err != nil {
		return false
	}
	return true
}
