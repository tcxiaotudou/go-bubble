package routes

import (
	"github.com/gin-gonic/gin"
	"go-bubble/controller"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 静态资源路径
	r.Static("/static", "static")
	// 模板文件路径
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1 todo
	v1Group := r.Group("v1")
	{
		// 新增待办事项
		v1Group.POST("/todo", controller.Create)

		// 查看所有代办事项
		v1Group.GET("/todo", controller.GetAll)

		// 修改某一个代办事项
		v1Group.PUT("/todo/:id", controller.UpdateByID)

		// 删除某一个代办事项
		v1Group.DELETE("/todo/:id", controller.DeleteByID)
	}
	return r
}
