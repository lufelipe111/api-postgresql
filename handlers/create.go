package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/study/api-postgresql/models"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error on decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var response map[string]any

	if err != nil {
		response = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error on inserting data: %v", err),
		}
	} else {
		response = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("To do inserted with success: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
