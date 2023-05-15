package model

type Todo struct {
	Id    uint   `json:"id"`
	Label string `json:"label"`
	Done  bool   `json:"done"`
}
