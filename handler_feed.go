package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Insider-Ekin-Aslan/Web-Application-RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handlerCreateFeed(writer http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(request.Body)

	params := parameters{}
	decoderError := decoder.Decode(&params)

	if decoderError != nil {
		respondError(writer, 400, fmt.Sprintf("JSON parsing error. ERROR: %v", decoderError))
		return
	}

	feed, feedError := config.Database.CreateFeed(request.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if feedError != nil {
		respondError(writer, 400, fmt.Sprintf("Couldn't create feed. ERROR: %v", feedError))
		return
	}

	respondJSON(writer, 201, databaseFeedToFeed(feed))
}

func (config *apiConfig) handlerGetFeeds(writer http.ResponseWriter, request *http.Request) {
	feed, feedError := config.Database.GetFeeds(request.Context())

	if feedError != nil {
		respondError(writer, 400, fmt.Sprintf("Couldn't get feeds. ERROR: %v", feedError))
		return
	}

	respondJSON(writer, 201, databaseFeedsToFeeds(feed))
}
