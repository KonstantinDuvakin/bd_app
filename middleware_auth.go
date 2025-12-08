package main

import (
	"net/http"

	"github.com/KonstantinDuvakin/bd_app/internal/auth"
	"github.com/KonstantinDuvakin/bd_app/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (ac *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, http.StatusForbidden, "no API key")
		}

		user, err := ac.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get user")
			return
		}

		handler(w, r, user)
	}
}
