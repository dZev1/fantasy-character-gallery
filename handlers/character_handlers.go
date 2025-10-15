package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dZev1/character-gallery/database"
	"github.com/dZev1/character-gallery/models"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	newCharacter := &models.Character{}

	err := json.NewDecoder(r.Body).Decode(newCharacter)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newCharacter.Name = strings.ToLower(newCharacter.Name)

	err = database.CreateCharacter(newCharacter)
	if err != nil {
		http.Error(w, "Could not create character", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*newCharacter)
}

func EditCharacter(w http.ResponseWriter, r *http.Request) {
	
}