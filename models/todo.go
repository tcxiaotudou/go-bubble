package models

import "go-bubble/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 增删改查
func Create(todo *Todo) (err error) {
	return dao.DB.Debug().Create(&todo).Error
}

func GetAll() (todoList []*Todo, err error) {
	return todoList, dao.DB.Debug().Find(&todoList).Error
}

func UpdateByID(todo *Todo) (err error) {
	return dao.DB.Debug().Model(&Todo{}).Where("id = ?", todo.ID).Update("status", todo.Status).Error
}

func DeleteByID(id int) (err error) {
	return dao.DB.Debug().Where("id = ?", id).Delete(Todo{}).Error
}
