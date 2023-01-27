package dto

import (
	"time"
	"zaizwk/ginessential/model"
)

type MsgDto struct {
	ID           uint      `json:"id"`
	FromUserName string    `json:"fromUserName"`
	ToUserName   string    `json:"toUserName"`
	CreatedAt    time.Time `json:"createdAt"`
	Content      string    `json:"content"`
}

func ToMsgDto(msg model.Message) MsgDto {
	return MsgDto{
		ID:           msg.ID,
		FromUserName: msg.FromUser.Name,
		ToUserName:   msg.ToUser.Name,
		CreatedAt:    msg.CreatedAt,
		Content:      msg.Content,
	}
}
