package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"zaizwk/ginessential/common"
	"zaizwk/ginessential/dto"
	"zaizwk/ginessential/model"
	"zaizwk/ginessential/response"
	"zaizwk/ginessential/util"
)

func ShowUserList(c *gin.Context) {
	DB := common.GetDB()
	var users []model.User
	DB.Find(&users)

	var userDTOs []dto.UserDto
	for _, u := range users {
		userDTOs = append(userDTOs, dto.ToUserDto(u))
	}
	response.Success(c, gin.H{"users": userDTOs}, "查询所有用户成功!")
}

func Register(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数 后台从表单获取数据
	var requestUser = model.User{}
	_ = c.Bind(&requestUser)
	name := requestUser.Name
	tel := requestUser.Tel
	pwd := requestUser.Pwd
	ps := requestUser.PS
	role := requestUser.Role
	// ApiPost
	//name := c.PostForm("name")
	//tel := c.PostForm("tel")
	//ps := c.PostForm("ps")

	// 数据验证
	if len(tel) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须是11位数并且不能以0开头")
		return
	}

	if len(pwd) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	// 如果账户没有传入，给一个10位随机账户
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// 手机号
	if isTelExist(DB, tel) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "该用户已经存在")
		return
	}
	// 创建用户
	hidePwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密出现错误")
		return
	}
	ps = "这个人很懒，什么都没有留下..."
	role = "worker"
	newUser := model.User{Name: name, Tel: tel, Pwd: string(hidePwd), PS: ps, Role: role}
	DB.Create(&newUser)
	// 结果
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Println(err)
		return
	}
	// 结果
	response.Success(c, gin.H{"token": token}, "注册成功")
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	// 前端获取参数
	var requestUser = model.User{}
	_ = c.Bind(&requestUser)
	tel := requestUser.Tel
	pwd := requestUser.Pwd

	// ApiPost
	//tel := c.PostForm("tel")
	//pwd := c.PostForm("pwd")
	// 数据验证
	if len(tel) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须是11位数并且不能以0开头")
		return
	}

	if len(pwd) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	// 手机
	if !isTelExist(DB, tel) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "该用户不存在")
		return
	}
	var user model.User
	DB.Where("tel = ?", tel).First(&user)
	// 密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Println(err)
		return
	}
	// 结果
	response.Success(c, gin.H{"token": token}, "登录成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			// 敏感信息的保护，如果只用 user，那么密码也会被返回
			"user": dto.ToUserDto(user.(model.User)),
		},
	})
}

func EditInfo(c *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	_ = c.Bind(&requestUser)
	name := requestUser.Name
	tel := requestUser.Tel
	ps := requestUser.PS
	role := requestUser.Role

	var user model.User
	DB.Where("tel = ?", tel).First(&user)
	DB.Model(&user).Updates(model.User{Name: name, PS: ps, Role: role})
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Println(err)
		return
	}
	response.Success(c, gin.H{"res": dto.ToUserDto(user), "token": token}, "信息修改成功!")
}

func RemoveUser(c *gin.Context) {
	DB := common.GetDB()
	tel := c.Param("tel")
	user := model.User{}
	DB.Where("tel = ?", tel).First(&user)
	DB.Where("user_tel = ?", tel).Delete(&model.Post{})
	DB.Unscoped().Delete(&model.User{}, user.ID)
	response.Success(c, gin.H{"flag": true, "tel": tel}, "删除成功!")
}

func isTelExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("tel = ?", tel).First(&user)
	return user.ID != 0
}
