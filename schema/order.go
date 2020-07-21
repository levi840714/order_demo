package schema

import "time"

type Order struct {
	ID        int       `gorm:"column:id;AUTO_INCREMENT"`
	AccountId string    `gorm:"column:accountId;type:varchar(12);FOREIGNKEY:;NOT NULL"`
	Goods     string    `gorm:"column:goods;type:varchar(20);NOT NULL"`
	Amount    float64   `gorm:"column:amount;type:decimal(12, 2);NOT NULL"`
	Status    string    `gorm:"column:status;type:enum('0', '1');NOT NULL"`
	OrderTime time.Time `gorm:"column:orderTime;type:datetime;NOT NULL"`
}

func (Order) TableName() string {
	return "order"
}
