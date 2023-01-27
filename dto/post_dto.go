package dto

import (
	"time"
	"zaizwk/ginessential/model"
)

type PostDto struct {
	Id       uint      `json:"id"`
	Tel      string    `json:"tel"`
	UserName string    `json:"userName"`
	CreateAt time.Time `json:"createdAt"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	PrizeCnt int64     `json:"prizeCnt"`
	IsTop    int8      `json:"top"`
}

func ToPostDto(post model.Post) PostDto {
	return PostDto{
		Id:       post.ID,
		Tel:      post.User.Tel,
		UserName: post.User.Name,
		CreateAt: post.CreatedAt,
		Title:    post.Title,
		Content:  post.Content,
		PrizeCnt: post.PrizeCnt,
		IsTop:    post.IsTop,
	}
}
