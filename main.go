package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
	"zaizwk/ginessential/common"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			panic("数据库关闭失败" + err.Error())
		}
	}(db)
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	_ = r.Run(":" + port)
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
