package data

import "myserver/pkg/dto"

var Todos []dto.TODO

func construct() {
	Todos = []dto.TODO{
		{ID: 1, Name: "Monday", IsDone: false},
		{ID: 2, Name: "Tuesday", IsDone: false},
		{ID: 3, Name: "Wednesday", IsDone: false},
		{ID: 4, Name: "Thursday", IsDone: false},
		{ID: 5, Name: "Friday", IsDone: false},
	}
}
