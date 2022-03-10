package routes

import (
	controllers "todo-list/example/controllers"
	middlewares "todo-list/example/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/auth", controllers.AuthHandler)

	router.GET("/users", controllers.UserIndex)

	router.POST("/user", controllers.UserStore)

	// router.PUT("/user/:id", controllers.UserUpdate)

	// router.DELETE("/user/:id", controllers.UserDestroy)

	router.GET("/todos", middlewares.JWTAuthMiddleware(), controllers.TodoIndex)
	router.POST("/todo", middlewares.JWTAuthMiddleware(), controllers.TodoStore)
	router.PATCH("/todo/:id", middlewares.JWTAuthMiddleware(), controllers.TodoUpdate)
	router.DELETE("/todo/:id", middlewares.JWTAuthMiddleware(), controllers.TodoDestory)

	return router
}
