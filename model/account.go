package model

import (
	"errors"
	"order_demo/lib"
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
	hash, _ := lib.HashPassword(password)
	createData := Account{Account: account, Password: hash, Status: "1", Balance: 0.0, CreateAt: time.Now()}
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

func CheckLogin(account string, password string) (bool, error) {
	var login Account
	if err := DB.Select("password").Where("account = ?", account).Find(&login).Error; err != nil {
		return false, errors.New("wrong Account")
	}
	check := lib.CheckPasswordHash(password, login.Password)
	if !check {
		return false, errors.New("wrong password")
	}
	return true, nil
}
