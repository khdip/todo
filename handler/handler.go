package handler

import (
	"html/template"
)

type ToDo struct {
	Task        string `json:"task"`
	IsCompleted bool   `json:"is_completed"`
}

type Handler struct {
	templates *template.Template
	ToDos     []ToDo
}

func GetHandler(todos []ToDo) *Handler {
	hand := &Handler{
		ToDos: todos,
	}
	hand.GetTemplate()
	return hand
}

func (h *Handler) GetTemplate() {
	h.templates = template.Must(template.ParseFiles("templates/create-todo.html", "templates/list-todo.html", "templates/edit-todo.html"))
}
