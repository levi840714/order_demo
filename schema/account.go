package schema

import "time"

type Account struct {
	ID       int       `gorm:"column:id;AUTO_INCREMENT"`
	Account  string    `gorm:"column:account;type:varchar(12);UNIQUE;NOT NULL"`
	Password string    `gorm:"column:password;type:varchar(60);NOT NULL"`
	Status   string    `gorm:"column:status;type:enum('0', '1');NOT NULL"`
	Balance  float64   `gorm:"column:balance;type:decimal(12, 2);NOT NULL"`
	CreateAt time.Time `gorm:"column:createAt;type:datetime;NOT NULL"`
}

func (Account) TableName() string {
	return "account"
}
