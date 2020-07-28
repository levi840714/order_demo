package schema

import "time"

type Order struct {
	ID        int       `gorm:"column:id;AUTO_INCREMENT"`
	AccountId int       `gorm:"column:accountId;int;NOT NULL"`
	GoodsId   int       `gorm:"column:goodsId;type:int;NOT NULL"`
	Status    string    `gorm:"column:status;type:enum('0', '1');NOT NULL"`
	OrderTime time.Time `gorm:"column:orderTime;type:datetime;NOT NULL"`
}

func (Order) TableName() string {
	return "order"
}
