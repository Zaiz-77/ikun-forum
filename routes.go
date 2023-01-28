package main

import (
	"github.com/gin-gonic/gin"
	"zaizwk/ginessential/controller"
	"zaizwk/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CROSMiddleware())
	// user
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.GET("api/auth/info/:tel", controller.ShowSpecificInfo)
	r.GET("api/auth/list", controller.ShowUserList)
	r.PUT("api/auth/update", controller.EditInfo)
	r.DELETE("api/auth/delete/:tel", controller.RemoveUser)

	// forum
	r.GET("api/auth/forum", controller.ShowPostList)
	r.POST("api/auth/forum/publish", middleware.AuthMiddleware(), controller.Publish)
	r.PUT("api/auth/forum/prize/:id", controller.Prize)
	r.PUT("api/auth/forum/top/:id", controller.TopSolve)
	r.DELETE("api/auth/forum/delete/:id", controller.RemovePost)

	// message
	r.GET("api/auth/msg/list", middleware.AuthMiddleware(), controller.ShowMsgList)
	r.POST("api/auth/msg/send/:id", middleware.AuthMiddleware(), controller.SendMsg)
	return r
}
