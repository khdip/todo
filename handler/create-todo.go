package handler

import (
	"net/http"
)

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "create-todo.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) StoreTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	myTask := r.FormValue("Task")
	if myTask == "" {
		http.Redirect(w, r, "/todo/create", http.StatusTemporaryRedirect)
		return
	}
	h.ToDos = append(h.ToDos, ToDo{Task: myTask})
	http.Redirect(w, r, "/todo", http.StatusTemporaryRedirect)
}

func (h *Handler) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	getTask := r.URL.Path[len("/todo/complete/"):]
	if getTask == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	for i, singleToDo := range h.ToDos {
		if singleToDo.Task == getTask {
			h.ToDos[i].IsCompleted = true
			break
		}
	}
	http.Redirect(w, r, "/todo", http.StatusTemporaryRedirect)
}

func (h *Handler) EditTodo(w http.ResponseWriter, r *http.Request) {
	getTask := r.URL.Path[len("/todo/edit/"):]
	if getTask == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}
	var newToDo ToDo
	for _, singleToDo := range h.ToDos {
		if singleToDo.Task == getTask {
			newToDo = singleToDo
			break
		}
	}
	err := h.templates.ExecuteTemplate(w, "edit-todo.html", newToDo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	getTask := r.URL.Path[len("/todo/update/"):]
	if getTask == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	myNewTask := r.FormValue("Task")
	if myNewTask == "" {
		http.Redirect(w, r, "/todo/edit", http.StatusTemporaryRedirect)
		return
	}

	for i, singleToDo := range h.ToDos {
		if singleToDo.Task == getTask {
			h.ToDos[i].Task = myNewTask
			break
		}
	}
	http.Redirect(w, r, "/todo", http.StatusTemporaryRedirect)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	getTask := r.URL.Path[len("/todo/delete/"):]
	if getTask == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	var todosAfterDelete []ToDo
	for _, singleToDo := range h.ToDos {
		if singleToDo.Task == getTask {
			continue
		}
		todosAfterDelete = append(todosAfterDelete, singleToDo)
	}
	h.ToDos = todosAfterDelete
	http.Redirect(w, r, "/todo", http.StatusTemporaryRedirect)
}
