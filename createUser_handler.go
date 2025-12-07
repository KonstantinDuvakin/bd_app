package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KonstantinDuvakin/bd_app/internal/database"
	"github.com/google/uuid"
)

func (ac *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type param struct {
		Name string `json:"name"`
	}
	decode := json.NewDecoder(r.Body)

	params := param{}
	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := ac.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid request payload")
		return
	}

	respondWithJSON(w, http.StatusCreated, dbUserToUser(user))
}
