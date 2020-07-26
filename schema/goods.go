package schema

import "time"

type Goods struct {
	ID       int       `gorm:"column:id;AUTO_INCREMENT"`
	Good     string    `gorm:"column:good;type:varchar(12);NOT NULL`
	Amount   float64   `gorm:"column:amount;type:decimal(12, 2);NOT NULL"`
	Status   string    `gorm:"column:status;type:enum('0', '1');NOT NULL"`
	CreateAt time.Time `gorm:"column:createAt;type:datetime;NOT NULL"`
}

func (Goods) TableName() string {
	return "goods"
}
