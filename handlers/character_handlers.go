package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dZev1/character-gallery/models/characters"
)

type CharacterHandler struct {
	Gallery characters.CharacterGallery
}

func (h *CharacterHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	newCharacter := &characters.Character{}

	err := json.NewDecoder(r.Body).Decode(newCharacter)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.Gallery.Create(newCharacter)
	if err != nil {
		http.Error(w, "Could not create character", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*newCharacter)
}

func (h *CharacterHandler) GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	chars, err := h.Gallery.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusFailedDependency)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chars)
}

func (h *CharacterHandler) GetCharacter(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	character, err := h.Gallery.Get(characters.ID(id))
	if err != nil {
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(character)
}

func (h *CharacterHandler) EditCharacter(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	characterToEdit := &characters.Character{}
	err = json.NewDecoder(r.Body).Decode(characterToEdit)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	characterToEdit.ID = characters.ID(id)
	characterToEdit.Stats.ID = characters.ID(id)
	characterToEdit.Customization.ID = characters.ID(id)

	err = h.Gallery.Edit(characterToEdit)
	if err != nil {
		http.Error(w, "Could not edit character", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characterToEdit)
}

func (h *CharacterHandler) DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Gallery.Remove(characters.ID(id))
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
