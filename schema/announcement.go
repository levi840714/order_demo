package schema

import "time"

type Announcement struct {
	ID      int       `gorm:"column:id;AUTO_INCREMENT"`
	Content string    `gorm:"column:content;type:varchar(80);NOT NULL"`
	Status  string    `gorm:"column:status;type:enum('0', '1');NOT NULL"`
	Date    time.Time `gorm:"column:date;type:date;UNIQUE;NOT NULL"`
}

func (Announcement) TableName() string {
	return "announcement"
}
