package services

import (
	"encoding/json"
	"log"
	"net/http"
	"todos/database"
	"todos/filters"
	"todos/models"
	"todos/utils"
)

// TodoList Retreive all todos
func TodoList(w http.ResponseWriter, r *http.Request) []byte {
	var todos []models.Todos
	db := database.DB.Model(todos)
	// db.Scopes(AmountGreaterThan1000).
	filters.Name(r, db)
	filters.Done(r, db)
	db.Order("name asc").Find(&todos)

	response := utils.WrapJSONResponse(todos)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return jsonResponse
}

// CreateTodo create todo
// todo := models.Todos{
// 	Name: "Todo 6",
// 	Done: false,
// }
func CreateTodo(w http.ResponseWriter, todo models.Todos) {

	if database.DB.Create(&todo).Error != nil {
		log.Panic("Unable to create todo: ")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
