package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KonstantinDuvakin/bd_app/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (ac *apiConfig) createFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type param struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decode := json.NewDecoder(r.Body)

	params := param{}
	err := decode.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feedFollow, err := ac.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't follow to a feed, %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, dbFeedFollowToFeedFollow(feedFollow))
}

func (ac *apiConfig) getFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := ac.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	respondWithJSON(w, http.StatusOK, dbFeedFollowsToFeedFollows(feedFollows))
}

func (ac *apiConfig) deleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse follow feed id")
		return
	}

	errDelete := ac.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if errDelete != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't delete the following feed")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
