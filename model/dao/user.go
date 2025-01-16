package dao

import "time"

type User struct {
	Id          int       `gorm:"column:id;type:int;primary_key"`
	FirstName   string    `gorm:"column:first_name;type:varchar(255)"`
	LastName    string    `gorm:"column:last_name;type:varchar(255)"`
	Email       string    `gorm:"column:email;type:varchar(255);not null;unique"`
	Age         int       `gorm:"column:age;type:int"`
	CreatedDate time.Time `gorm:"column:created_date;type:timestamptz;not null"`
}
