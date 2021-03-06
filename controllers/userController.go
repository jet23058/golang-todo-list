package controllers

import (
	"net/http"
	"strconv"
	"time"
	model "todo-list/example/models"
	"todo-list/example/responses"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//列表数据
func UserIndex(c *gin.Context) {
	var user model.User
	user.Name = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Users()

	if err != nil {
		panic(c)
	}

	responses.Response(c, &responses.Success{
		Data: result,
	})
}

type UserInsert struct {
	Name   string `json:name`
	Password string `json:password`
}

//添加数据
func UserStore(c *gin.Context) {
	var user model.User

	json := UserInsert{}
	c.BindJSON(&json)

	user.Uuid = uuid.New()
	user.Name = json.Name
	user.Password = json.Password
	user.Status = "registered"
	user.Created_at = time.Now()
	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    id,
	})
}

//修改数据
func UserUpdate(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改成功",
	})
}

//删除数据
func UserDestroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除成功",
	})
}
