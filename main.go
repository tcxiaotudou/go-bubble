package main

import (
	"go-bubble/dao"
	"go-bubble/models"
	"go-bubble/routes"
)

func main() {
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	// 创建表
	dao.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()
	r.Run()
}
