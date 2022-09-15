package handlers

import (
	"encoding/json"
	"github.com/study/api-postgresql/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error on retrieve registers: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
