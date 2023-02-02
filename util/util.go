package util

import (
	"github.com/jinzhu/gorm"
	"zaizwk/ginessential/model"
)

func IsTelExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("tel = ?", tel).First(&user)
	return user.ID != 0
}

func IsNameExist(db *gorm.DB, name string) bool {
	var user model.User
	db.Where("name = ?", name).First(&user)
	return user.ID != 0
}
