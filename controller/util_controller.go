package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zaizwk/ginessential/common"
	"zaizwk/ginessential/dto"
	"zaizwk/ginessential/model"
	"zaizwk/ginessential/response"
)

func ShowSpecificInfo(c *gin.Context) {
	DB := common.GetDB()
	tel := c.Param("tel")
	var user model.User
	DB.Where("tel = ?", tel).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "该用户已经离开了本站...")
		return
	}
	response.Success(c, gin.H{"res": dto.ToUserDto(user)}, "查询信息成功!")
}
