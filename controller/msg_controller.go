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
	DB.Preload("ToUser").Preload("FromUser").Where("to_user_tel = ?", now.Tel).Order("created_at DESC").Find(&msg)
	var msgDTO []dto.MsgDto
	for _, m := range msg {
		msgDTO = append(msgDTO, dto.ToMsgDto(m))
	}
	response.Success(c, gin.H{"allMsg": msgDTO}, "查询所有消息成功!")
}
