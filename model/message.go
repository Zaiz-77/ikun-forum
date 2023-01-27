package model

import "time"

type Message struct {
	ID          uint      `gorm:"primary_key"`
	CreatedAt   time.Time `json:"createdAt"`
	FromUser    User      `gorm:"ForeignKey:FromUserTel"`
	ToUser      User      `gorm:"ForeignKey:ToUserTel"`
	FromUserTel string    `gorm:"varchar(11);not null"`
	ToUserTel   string    `gorm:"varchar(11);not null"`
	Content     string    `gorm:"not null"`
}
