package dao

import "time"

type User struct {
	Id          int       `gorm:"type:int;primary_key"`
	FirstName   string    `gorm:"type:varchar(255)"`
	LastName    string    `gorm:"type:varchar(255)"`
	Email       string    `gorm:"type:varchar(255);not null;unique"`
	Age         int       `gorm:"type:int"`
	CreatedDate time.Time `gorm:"type:timestamp;not null"`
}
