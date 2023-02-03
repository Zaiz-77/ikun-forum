package controller

import (
	"github.com/gin-gonic/gin"
	"zaizwk/ginessential/common"
	"zaizwk/ginessential/dto"
	"zaizwk/ginessential/model"
	"zaizwk/ginessential/response"
)

func ReadMsg(c *gin.Context) {
	DB := common.GetDB()
	var msg []model.Message
	_ = c.Bind(&msg)
	for _, m := range msg {
		DB.Model(&m).Update("is_read", true)
	}
	response.Success(c, gin.H{"readMsg": msg, "cnt": len(msg)}, "阅读所有未读消息成功!")
}

func ShowMsgList(c *gin.Context) {
	DB := common.GetDB()
	t, _ := c.Get("user")
	now := t.(model.User)

	var msg []model.Message
	DB.Preload("ToUser").Preload("FromUser").
		Where("to_user_tel = ? AND is_read = ?", now.Tel, 0).
		Order("created_at ASC").Find(&msg)

	mp := make(map[string][]dto.MsgDto)
	for _, m := range msg {
		mp[m.FromUser.Name] = append(mp[m.FromUser.Name], dto.ToMsgDto(m))
	}
	response.Success(c, gin.H{"allMsg": mp, "cnt": len(msg)}, "查询所有未读消息成功!")
}

func ShowChatFace(c *gin.Context) {
	DB := common.GetDB()
	t, _ := c.Get("user")
	now := t.(model.User)

	var fromUser model.User
	DB.Where("name = ?", c.Param("name")).First(&fromUser)

	var (
		msg    []model.Message
		msgDTO []dto.MsgDto
	)
	DB.Preload("ToUser").Preload("FromUser").
		Where("to_user_tel = ? AND from_user_tel = ?", now.Tel, fromUser.Tel).
		Or("to_user_tel = ? AND from_user_tel = ?", fromUser.Tel, now.Tel).
		Order("created_at ASC").Find(&msg)
	for _, m := range msg {
		cur := dto.ToMsgDto(m)
		if now.Tel == m.FromUserTel {
			cur.IsUser = true
		}
		msgDTO = append(msgDTO, cur)
	}

	response.Success(c, gin.H{"allMsg": msgDTO}, "查阅消息记录成功!")
}

func SendMsg(c *gin.Context) {
	DB := common.GetDB()
	t, _ := c.Get("user")
	now := t.(model.User)

	var toUser model.User
	DB.Where("name = ?", c.Param("name")).First(&toUser)

	var req model.Message
	_ = c.Bind(&req)
	content := req.Content
	message := model.Message{
		FromUser: now,
		ToUser:   toUser,
		Content:  content,
		IsRead:   false,
	}
	DB.Save(&message)
	response.Success(c, gin.H{"msg": dto.ToMsgDto(message)}, "发送消息成功!")
}
