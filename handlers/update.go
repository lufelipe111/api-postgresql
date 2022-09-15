package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/study/api-postgresql/models"
	"log"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error on parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error on decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Error on updating todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Warning: %d registers were updated", rows)
	}

	response := map[string]any{
		"Error":   false,
		"Message": "Data updated with success",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
