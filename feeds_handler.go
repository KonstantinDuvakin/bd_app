package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KonstantinDuvakin/bd_app/internal/database"
	"github.com/google/uuid"
)

func (ac *apiConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type param struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decode := json.NewDecoder(r.Body)

	params := param{}
	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := ac.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		Owner:     user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create a feed")
		return
	}

	respondWithJSON(w, http.StatusCreated, dbFeedToFeed(feed))
}

func (ac *apiConfig) getAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := ac.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't return feeds")
		return
	}

	respondWithJSON(w, http.StatusCreated, dbFeedsToFeeds(feeds))
}
