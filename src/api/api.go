package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/alexbvart/go-pq/src/helpers"
	"github.com/alexbvart/go-pq/src/models"
	"github.com/gorilla/mux"
)

type Data struct {
	Success bool          `json: "success`
	Data    []models.Todo `json: "data"`
	Errors  []string      `json:"errors"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	bodyTodo, success := helpers.DecodeBody(r)
	if success != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)

	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid description")

		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	todo, success := models.Insert(bodyTodo.Description)
	if success != true {
		data.Errors = append(data.Errors, "could not create todo")
	}

	data.Success = true
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data Data
	var todo models.Todo
	var success bool

	todo, success = models.Get(id)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "todo not found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	data.Success = true
	data.Data = append(data.Data, todo)
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo = models.GetTodos()

	var data = Data{true, todos, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data = Data{Errors: make([]string, 0)}

	todo, success := models.Delete(id)

	if success != true {
		data.Errors = append(data.Errors, "could not delete todo")
	}

	data.Success = true
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	bodyTodo, success := helpers.DecodeBody(r)
	if success != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}
	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	todo, success := models.Update(id, bodyTodo.Description)
	if success != true {
		data.Errors = append(data.Errors, "could not update todo")
	}

	data.Success = true
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
