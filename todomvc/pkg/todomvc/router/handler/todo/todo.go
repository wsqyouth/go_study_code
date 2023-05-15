package todo

import (
	"net/http"
	"todomvc/pkg/todomvc/model"

	"github.com/gin-gonic/gin"
)

type TodoRequest struct {
	Label string `json:"label"`
}

func List(c *gin.Context) {
	var todos []model.Todo
	model.GetDB().Where("").Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func Get(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	model.GetDB().First(&todo, id)
	c.JSON(http.StatusOK, todo)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	model.GetDB().Delete(&model.Todo{}, id)
	c.JSON(http.StatusOK, "")
}

func Create(c *gin.Context) {
	todoRequest := TodoRequest{}
	c.BindJSON(&todoRequest)
	var todo = model.Todo{Label: todoRequest.Label}
	model.GetDB().Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func Update(c *gin.Context) {
	todoRequest := TodoRequest{}
	c.BindJSON(&todoRequest)
	id := c.Param("id")
	var todo model.Todo
	model.GetDB().First(&todo, id)
	todo.Label = todoRequest.Label
	model.GetDB().Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func Toggle(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	model.GetDB().First(&todo, id)
	todo.Done = !todo.Done
	model.GetDB().Save(&todo)
	c.JSON(http.StatusOK, todo)
}
