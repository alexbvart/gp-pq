package main

import (
	"fmt"
	"net/http"

	"github.com/alexbvart/go-pq/src/api"
	"github.com/gorilla/mux"
)

func main() {
	var port string = "8080"
	router := mux.NewRouter()
	apirouter := router.PathPrefix("/api/").Subrouter()
	apirouter.HandleFunc("/todos", api.GetTodos).Methods("GET")
	apirouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")
	apirouter.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	apirouter.HandleFunc("/todos/{id}", api.UpdateTodo).Methods("PUT")
	apirouter.HandleFunc("/todos/{id}", api.DeleteTodo).Methods("DELETE")

	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)
}
