package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Insider-Ekin-Aslan/Web-Application-RSS-Aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (config *apiConfig) handlerCreateFeedFollow(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(request.Body)

	params := parameters{}
	decoderError := decoder.Decode(&params)

	if decoderError != nil {
		respondError(writer, 400, fmt.Sprintf("JSON parsing error. ERROR: %v", decoderError))
		return
	}

	feedFollow, feedFollowError := config.Database.CreateFeedFollow(request.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if feedFollowError != nil {
		respondError(writer, 400, fmt.Sprintf("Couldn't create feed follow. ERROR: %v", feedFollowError))
		return
	}

	respondJSON(writer, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (config *apiConfig) handlerGetFeedFollows(writer http.ResponseWriter, request *http.Request, user database.User) {
	feedFollows, feedFollowError := config.Database.GetFeedFollows(request.Context(), user.ID)

	if feedFollowError != nil {
		respondError(writer, 400, fmt.Sprintf("Couldn't get feed follow. ERROR: %v", feedFollowError))
		return
	}

	respondJSON(writer, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (config *apiConfig) handlerDeleteFeedFollow(writer http.ResponseWriter, request *http.Request, user database.User) {
	id, idError := uuid.Parse(chi.URLParam(request, "id"))

	if idError != nil {
		respondError(writer, 400, fmt.Sprintf("Couldn't parse feed follow ID. ERROR: %v", idError))
		return
	}

	deleteError := config.Database.DeleteFeedFollows(request.Context(), database.DeleteFeedFollowsParams{
		ID:     id,
		UserID: user.ID,
	})

	if deleteError != nil {
		respondError(writer, 400, fmt.Sprintf("Couldn't delete feed follow. ERROR: %v", deleteError))
		return
	}

	respondJSON(writer, 200, struct{}{})
}
