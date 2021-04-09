package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"todos/database"
	"todos/models"
	"todos/services"
	"todos/utils"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// TodosList godoc
func TodosList(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)
	response := services.TodoList(w, r)
	w.Write(response)
}

// CreateTodo godoc
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todos

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	services.CreateTodo(w, todo)
}

// ShowTodo retreive a single todo via id
func ShowTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if !utils.IsValidUUID(id) {
		http.Error(w, "Invalid uuid", http.StatusBadRequest)
		return
	}
	var todo = models.Todos{}
	database.DB.Model(todo).Where("id = ?", id).Find(&todo)
	jsonResponse, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

// UpdateTodo update a single todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// enableCors(&w)
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// if r.Method == "OPTIONS" {
	// 	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	// }

	vars := mux.Vars(r)
	id := vars["id"]
	if !utils.IsValidUUID(id) {
		http.Error(w, "Invalid uuid", http.StatusBadRequest)
		return
	}

	var todo models.Todos

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Please send a valid json body", http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	database.DB.Model(&todo).Where("id = ?", id).Updates(&todo)
}

// DeleteTodo godoc
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if !utils.IsValidUUID(id) {
		http.Error(w, "Invalid uuid", http.StatusBadRequest)
		return
	}
	var todo models.Todos
	database.DB.Model(&todo).Where("id = ?", id).Update("deleted", true)
}
