package main

import (
	"fmt"
	"log"
	"net/http"
	"practice/todo/handler"
)

func main() {
	myTodos := []handler.ToDo{
		{
			Task:        "Learn Go",
			IsCompleted: true,
		},
		{
			Task:        "Get a Project",
			IsCompleted: false,
		},
		{
			Task:        "Become a Pro",
			IsCompleted: false,
		},
	}
	h := handler.GetHandler(myTodos)
	http.HandleFunc("/todo", h.GetTodo)
	http.HandleFunc("/todo/create", h.CreateTodo)
	http.HandleFunc("/todo/store", h.StoreTodo)
	http.HandleFunc("/todo/complete/", h.CompleteTodo)
	http.HandleFunc("/todo/edit/", h.EditTodo)
	http.HandleFunc("/todo/Update/", h.UpdateTodo)
	http.HandleFunc("/todo/delete/", h.DeleteTodo)
	fmt.Println("Server Starting...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server Not Found", err)
	}
}
