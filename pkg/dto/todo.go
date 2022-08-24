package dto

type TODO struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"isdone"`
}
