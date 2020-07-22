package model

import (
	"errors"
	"order_demo/lib/hash"
	"time"
)

const (
	StatusStop = "0"
	StatusOK   = "1"
)

type Account struct {
	ID       int       `gorm:"column:id"`
	Account  string    `gorm:"column:account"`
	Password string    `gorm:"column:password"`
	Role     int       `gorm:"column:role"`
	Status   string    `gorm:"column:status"`
	Balance  float64   `gorm:"column:balance"`
	CreateAt time.Time `gorm:"column:createAt"`
}

func (Account) TableName() string {
	return "account"
}

func RegisterAccount(account string, password string) (bool, error) {
	hash, _ := hash.HashPassword(password)
	createData := Account{Account: account, Password: hash, Status: "1", Balance: 0.0, CreateAt: time.Now()}
	if err := DB.Create(&createData).Error; err != nil {
		return false, err
	}
	return true, nil

}

func GetAccount(accountId int) (*Account, error) {
	var account Account
	if err := DB.Where("id = ?", accountId).Find(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func CheckLogin(account string, password string) (*Account, error) {
	var accountData Account
	if err := DB.Select("id, password, role").Where("account = ? AND status = ?", account, StatusOK).Find(&accountData).Error; err != nil {
		return nil, errors.New("wrong Account")
	}
	check := hash.CheckPasswordHash(password, accountData.Password)
	if !check {
		return nil, errors.New("wrong password")
	}
	return &accountData, nil
}
