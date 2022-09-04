package main

import (
	"github.com/labstack/echo/v4"

	"myserver/pkg/handler"
)

func main() {
	e := echo.New()
	e.GET("/todo", handler.ListTodo)

	e.POST("/todo", handler.CreateTodo)

	e.GET("/todo/:id", handler.GetTodoByID)

	e.DELETE("/todo/:id", handler.DeleteTodo)

	e.PUT("/todo/:id", handler.EditTodo)

	e.Logger.Fatal(e.Start(":3000"))
}
