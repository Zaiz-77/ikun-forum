package main

import (
	"github.com/gin-gonic/gin"
	"zaizwk/ginessential/controller"
	"zaizwk/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CROSMiddleware())
	// user
	r.POST("/api/auth/user/register", controller.Register)
	r.POST("/api/auth/user/login", controller.Login)
	r.GET("api/auth/user/info", middleware.AuthMiddleware(), controller.Info)
	r.GET("api/auth/user/info/:tel", controller.ShowSpecificInfo)
	r.GET("api/auth/user/list", controller.ShowUserList)
	r.PUT("api/auth/user/update", controller.EditInfo)
	r.DELETE("api/auth/user/delete/:tel", controller.RemoveUser)

	// forum
	r.GET("api/auth/forum", controller.ShowPostList)
	r.POST("api/auth/forum/publish", middleware.AuthMiddleware(), controller.Publish)
	r.PUT("api/auth/forum/prize/:id", controller.Prize)
	r.PUT("api/auth/forum/top/:id", controller.TopSolve)
	r.DELETE("api/auth/forum/delete/:id", controller.RemovePost)

	// message
	r.GET("api/auth/msg/list", middleware.AuthMiddleware(), controller.ShowMsgList)
	r.POST("api/auth/msg/send/:name", middleware.AuthMiddleware(), controller.SendMsg)
	r.PUT("api/auth/msg/read", controller.ReadMsg)
	r.GET("api/auth/msg/chat/:name", middleware.AuthMiddleware(), controller.ShowChatFace)

	// friend
	r.GET("api/auth/friend/list", middleware.AuthMiddleware(), controller.FriendList)
	return r
}
