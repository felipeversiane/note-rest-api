package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/felipeversiane/note-rest-api/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.List()

	if err != nil {
		log.Printf("%v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
