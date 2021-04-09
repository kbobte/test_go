package router

import (
	"todos/handlers"
	"todos/uploadFileAPI"

	"github.com/gorilla/mux"
)

// GetRouter godoc
func GetRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.TodosList).Methods("GET").Name("Retreive list of todos")
	// r.HandleFunc("/todos", handlers.TodosDoneList).Methods("GET")
	// r.HandleFunc("/todos", handlers.TodosNotDoneList).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.ShowTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	r.HandleFunc("/todos/{id}/upload", uploadFileAPI.UploadFile).Methods("POST")
	return r
}
