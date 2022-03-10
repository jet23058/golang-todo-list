package main

import (
	_ "todo-list/example/database"
	orm "todo-list/example/database"
	router "todo-list/example/routes"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":3388")
}
