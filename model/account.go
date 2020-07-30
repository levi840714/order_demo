package model

import (
	"errors"
	"order_demo/lib/compute"
	"order_demo/lib/hash"
	"time"
)

const (
	AccountStatusFail = "0"
	AccountStatusOK   = "1"
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
	createData := Account{Account: account, Password: hash, Status: AccountStatusOK, Role: 2, Balance: 0.0, CreateAt: time.Now()}
	if err := DB.Create(&createData).Error; err != nil {
		return false, err
	}
	return true, nil

}

func GetAccountInfo(accountId int) (*Account, error) {
	var account Account
	if err := DB.Where("id = ?", accountId).Find(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func CheckLogin(account string, password string) (*Account, error) {
	var accountData Account
	if err := DB.Select("id, password, role").Where("account = ? AND status = ?", account, AccountStatusOK).Find(&accountData).Error; err != nil {
		return nil, errors.New("wrong Account")
	}
	check := hash.CheckPasswordHash(password, accountData.Password)
	if !check {
		return nil, errors.New("wrong password")
	}
	return &accountData, nil
}

func UpdateBalance(accountId int, amount float64) (float64, error) {
	var accountData Account
	if err := TX.Set("gorm:query_option", "FOR UPDATE").Where("id = ? AND status = ?", accountId, AccountStatusOK).Find(&accountData).Error; err != nil {
		return 0, err
	}
	accountData.Balance = compute.Add(accountData.Balance, amount)
	if err := TX.Save(&accountData).Error; err != nil {
		return 0, err
	}
	return accountData.Balance, nil
}
