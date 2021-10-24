package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/alexbvart/go-pq/src/models"
)

func DecodeBody(r *http.Request) (models.Todo, bool) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		return models.Todo{}, false
	}
	return todo, true
}

func IsValidDescription(description string) bool {
	desc := strings.TrimSpace(description)
	if len(desc) == 0 {
		return false
	}
	return true
}
