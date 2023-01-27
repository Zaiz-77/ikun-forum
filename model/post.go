package model

import (
	"time"
)

type Post struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User   `gorm:"ForeignKey:UserTel"`
	UserName  string `gorm:"varchar(255);not null"`
	UserTel   string `gorm:"varchar(11);not null"`
	Title     string `gorm:"varchar(20);not null"`
	Content   string `gorm:"not null"`
	PrizeCnt  int64
	IsTop     int8 `gorm:"default:0"`
}
