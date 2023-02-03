package controller

import (
	"github.com/gin-gonic/gin"
	"zaizwk/ginessential/common"
	"zaizwk/ginessential/model"
	"zaizwk/ginessential/response"
)

func FriendList(c *gin.Context) {
	DB := common.GetDB()
	t, _ := c.Get("user")
	now := t.(model.User)

	var msg []model.Message
	var friend []string
	mp := make(map[string]bool)
	DB.Preload("ToUser").Preload("FromUser").
		Where("to_user_tel = ?", now.Tel).
		Or("from_user_tel = ?", now.Tel).
		Order("created_at ASC").Find(&msg)
	for _, m := range msg {
		if m.FromUserTel == now.Tel {
			mp[m.ToUser.Name] = true
		} else {
			mp[m.FromUser.Name] = true
		}
	}
	for f := range mp {
		friend = append(friend, f)
	}
	response.Success(c, gin.H{"allFriends": friend}, "查询所有朋友成功!")
}
