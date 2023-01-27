package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"varchar(20);not null"`
	Tel       string `gorm:"varchar(11);not null"`
	Pwd       string `gorm:"size:255;not null"`
	PS        string `gorm:"varchar(255)"`
	Role      string `gorm:"varchar(10);not null"`
}
