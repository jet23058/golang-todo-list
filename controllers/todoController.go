package controllers

import (
	"strconv"
	"time"
	"todo-list/example/models"
	model "todo-list/example/models"
	"todo-list/example/responses"

	orm "todo-list/example/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResponseTodo struct {
	Id int64
}

func TodoIndex(c *gin.Context) {
	var todos []model.Todo

	containDelete := c.DefaultQuery("contain_delete", "0")

	if containDelete == "1" {
		if err := orm.Eloquent.Unscoped().Find(&todos).Error; err != nil {
			panic(c)
		}
	} else {
		if err := orm.Eloquent.Find(&todos).Error; err != nil {
			panic(c)
		}
	}

	responses.Response(c, &responses.Success{
		Data: todos,
	})
}

type TodoInsert struct {
	Title   string `json:title`
	Content string `json:content`
	Status  string `json:status`
}

func TodoStore(c *gin.Context) {
	var todos model.Todo

	json := TodoInsert{}

	c.BindJSON(&json)

	userId, _ := c.Get("userId")

	todos.Uuid = uuid.New()
	todos.Title = json.Title
	todos.Content = json.Content
	todos.Status = models.TodoStauts(json.Status)
	todos.CreatedBy = userId.(int64)
	todos.CreatedAt = time.Now()
	result := orm.Eloquent.Create(&todos)

	if result.Error != nil {
		panic(c)
	}

	responses.Response(c, &responses.Success{
		Data: todos,
	})
}

type Update struct {
	Title   string `json:title`
	Content string `json:content`
	Status  string `json:status`
}

func TodoUpdate(c *gin.Context) {
	var todos, targetTodo model.Todo

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		panic(c)
	}

	json := Update{}
	c.BindJSON(&json)

	userId, _ := c.Get("userId")

	todos.Title = json.Title
	todos.Content = json.Content
	todos.Status = models.TodoStauts(json.Status)
	todos.UpdatedBy = userId.(int64)
	todos.UpdatedAt = time.Now()

	if err := orm.Eloquent.First(&targetTodo, id).Error; err != nil {
		panic(c)
	}

	if err := orm.Eloquent.Model(&targetTodo).Updates(&todos).Error; err != nil {
		panic(c)
	}

	responses.Response(c, &responses.Success{
		Data: targetTodo,
	})
}

func TodoDestory(c *gin.Context) {
	var todos model.Todo

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		panic(c)
	}

	if err := orm.Eloquent.First(&todos, id).Error; err != nil {
		panic(c)
	}

	userId, _ := c.Get("userId")
	todos.DeletedBy = userId.(int64)

	if err := orm.Eloquent.Model(&todos).Updates(&todos).Error; err != nil {
		panic(c)
	}

	if err := orm.Eloquent.Delete(&todos).Error; err != nil {
		panic(c)
	}

	responses.Response(c, &responses.Success{
		Data: todos,
	})
}
