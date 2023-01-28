package controller

import (
	"github.com/gin-gonic/gin"
	"zaizwk/ginessential/common"
	"zaizwk/ginessential/dto"
	"zaizwk/ginessential/model"
	"zaizwk/ginessential/response"
)

func ShowMsgList(c *gin.Context) {
	DB := common.GetDB()
	t, _ := c.Get("user")
	now := t.(model.User)
	var msg []model.Message
	DB.Preload("ToUser").Preload("FromUser").
		Where("to_user_tel = ?", now.Tel).
		Order("created_at DESC").Find(&msg)

	mp := make(map[string][]dto.MsgDto)
	for _, m := range msg {
		mp[m.FromUser.Name] = append(mp[m.FromUser.Name], dto.ToMsgDto(m))
	}
	response.Success(c, gin.H{"allMsg": mp, "cnt": len(msg)}, "查询所有消息成功!")
}

func SendMsg(c *gin.Context) {
	DB := common.GetDB()
	t, _ := c.Get("user")
	now := t.(model.User)

	var toUser model.User
	DB.First(&toUser, c.Param("id"))

	content := c.PostForm("content")
	message := model.Message{
		FromUser: now,
		ToUser:   toUser,
		Content:  content,
	}
	DB.Save(&message)
	response.Success(c, gin.H{"msg": dto.ToMsgDto(message)}, "发送消息成功!")
}
