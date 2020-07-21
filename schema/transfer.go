package schema

import "time"

type Transfer struct {
	ID           int       `gorm:"column:id;AUTO_INCREMENT"`
	AccountId    string    `gorm:"column:accountId;type:varchar(12);NOT NULL"`
	Amount       float64   `gorm:"column:amount;type:decimal(12, 2);NOT NULL"`
	TransferTime time.Time `gorm:"column:transferTime;type:datetime;NOT NULL"`
}

func (Transfer) TableName() string {
	return "transfer"
}
