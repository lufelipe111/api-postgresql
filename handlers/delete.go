package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/study/api-postgresql/models"
	"log"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error on parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Error on deleting todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Warning: %d registers were deleted", rows)
	}

	response := map[string]any{
		"Error":   false,
		"Message": "Data deleted with success",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
