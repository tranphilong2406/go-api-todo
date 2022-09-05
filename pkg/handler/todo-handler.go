package handler

import (
	"myserver/pkg/data"
	"myserver/pkg/dto"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ListTodo(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, data.Todos, "  ")
}

func GetTodoByID(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		return c.String(http.StatusBadRequest, "Sai kieu du lieu !")
	}

	for i := 0; i < len(data.Todos); i++ {

		if data.Todos[i].ID == id {
			td := data.Todos[i]
			return c.JSONPretty(http.StatusOK, td, "  ")
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Khong tim thay id nay")

}

func CreateTodo(c echo.Context) (err error) {
	u := new(dto.TODO)

	if err = c.Bind(u); err != nil {
		return
	}

	todo := dto.TODO{
		ID:     AutoID(data.Todos),
		Name:   u.Name,
		IsDone: false,
	}

	data.Todos = append(data.Todos, todo)

	return c.String(http.StatusOK, "Them thanh cong!")
}

func EditTodo(c echo.Context) (e error) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return c.String(http.StatusOK, "Sai kieu du lieu!")
	}

	todo := new(dto.TODO)

	if e = c.Bind(todo); e != nil {
		return
	}

	for i, v := range data.Todos {
		if v.ID == id {
			data.Todos[i].Name = todo.Name
			return c.String(http.StatusOK, "Sua thanh cong !")
		}
	}
	return c.String(http.StatusNotFound, "Khong the tim thay id nay!")

}

func DeleteTodo(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		return c.String(http.StatusNoContent, "Sai kieu du lieu!")
	}

	for i, todo := range data.Todos {
		if todo.ID == id {
			data.Todos = append(data.Todos[:i], data.Todos[i+1:]...)
			return c.String(http.StatusOK, "Xoa thanh cong!")
		}
	}
	return c.String(http.StatusNotFound, "Khong the tim thay id nay!")

}

func CheckTodo(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return c.String(http.StatusBadRequest, "Sai kieu du lieu!")
	}

	for i, v := range data.Todos {
		if v.ID == id {
			data.Todos[i].IsDone = true
			return c.String(http.StatusOK, "Cap nhat thanh cong!")
		}
	}

	return c.String(http.StatusNotFound, "Khong tim thay id nay!")
}

func AutoID(todos []dto.TODO) int {
	var id int = 0

	for _, v := range todos {
		if v.ID > id {
			id = v.ID
		}
	}
	return id + 1
}
