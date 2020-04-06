package controller

import (
	"github.com/gin-gonic/gin"
	"go-bubble/models"
	"net/http"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Create(c *gin.Context) {
	// 绑定参数
	var todo models.Todo
	c.BindJSON(&todo)
	// 保存并响应
	if err := models.Create(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
	}
	c.JSON(http.StatusOK, todo)
}

func GetAll(c *gin.Context) {
	todoList, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, todoList)
}

func UpdateByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效id"})
	}
	var todo models.Todo
	c.BindJSON(&todo)
	todo.ID, _ = strconv.Atoi(id)
	err := models.UpdateByID(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func DeleteByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效id"})
	}
	todoID, _ := strconv.Atoi(id)
	if err := models.DeleteByID(todoID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
