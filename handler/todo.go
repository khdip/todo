package handler

import (
	"net/http"
)

type ListToDo struct {
	ToDo_list []ToDo
}

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	lt := ListToDo{ToDo_list: h.ToDos}
	err := h.templates.ExecuteTemplate(w, "list-todo.html", lt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
