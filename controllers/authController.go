package controllers

import (
	"net/http"
	orm "todo-list/example/database"
	middleware "todo-list/example/middlewares"
	model "todo-list/example/models"

	"github.com/gin-gonic/gin"
)

type AuthForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func AuthHandler(c *gin.Context) {
	var user model.User

	json := AuthForm{}

	c.BindJSON(&json)

	err := orm.Eloquent.Where("name = ?", json.Name).Where("password = ? ", json.Password).First(&user).Error

	if err != nil {
		panic(c)
	}

	tokenString, _ := middleware.GenToken(user)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"token": tokenString},
	})
}
