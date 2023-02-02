package controller

import (
	"github.com/gin-gonic/gin"
	"zaizwk/ginessential/common"
	"zaizwk/ginessential/dto"
	"zaizwk/ginessential/model"
	"zaizwk/ginessential/response"
)

func ShowPostList(c *gin.Context) {
	DB := common.GetDB()
	var posts []model.Post
	DB.Preload("User").Order("is_top DESC, created_at DESC").Find(&posts)

	var postDTOs []dto.PostDto
	for _, post := range posts {
		postDTOs = append(postDTOs, dto.ToPostDto(post))
	}
	response.Success(c, gin.H{"res": postDTOs}, "查询所有帖子成功!")
}

func Publish(c *gin.Context) {
	DB := common.GetDB()
	var req model.Post
	_ = c.Bind(&req)
	title := req.Title
	content := req.Content

	temp, _ := c.Get("user")
	var now model.User
	now = temp.(model.User)
	post := model.Post{
		User:     now,
		Title:    title,
		Content:  content,
		PrizeCnt: 0,
	}
	DB.Create(&post)
	response.Success(c, gin.H{"post": dto.ToPostDto(post)}, "发表成功!")
}

func Prize(c *gin.Context) {
	DB := common.GetDB()
	// 前端传
	id := c.Param("id")

	// ApiPost
	//id := c.PostForm("id")
	post := model.Post{}
	DB.First(&post, id)
	DB.Model(&post).Where("id = ?", id).Update("prizeCnt", post.PrizeCnt+1)

	response.Success(c, gin.H{"prize_which": id, "likes": post.PrizeCnt}, "点赞成功!")
}

func TopSolve(c *gin.Context) {
	DB := common.GetDB()
	id := c.Param("id")
	var now model.Post
	var msg string
	DB.First(&now, id)
	if now.IsTop == 1 {
		now.IsTop = 0
		msg = "取消置顶成功!"
	} else {
		now.IsTop = 1
		msg = "置顶成功!"
	}
	DB.Save(&now)
	response.Success(c, gin.H{"post": dto.ToPostDto(now)}, msg)
}

func RemovePost(c *gin.Context) {
	DB := common.GetDB()
	id := c.Param("id")
	post := model.Post{}
	DB.First(&post, id)
	DB.Unscoped().Delete(&model.Post{}, id)
	response.Success(c, gin.H{"flag": true, "id": id}, "删除成功!")
}
