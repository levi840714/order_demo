package schema

import "time"

type Account struct {
	ID       int       `gorm:"column:id;AUTO_INCREMENT"`
	Account  string    `gorm:"column:account;type:varchar(12);UNIQUE;NOT NULL"`
	Password string    `gorm:"column:password;type:varchar(80);NOT NULL"`
	Role     int       `gorm:"column:role;type:int;default:2;comment:'1=>admin,2=>member'"`
	Status   string    `gorm:"column:status;type:enum('0', '1');default:'1'"`
	Balance  float64   `gorm:"column:balance;type:decimal(12, 2);default:0"`
	CreateAt time.Time `gorm:"column:createAt;type:datetime;NOT NULL"`
}

func (Account) TableName() string {
	return "account"
}
