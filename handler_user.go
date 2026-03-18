package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Insider-Ekin-Aslan/Web-Application-RSS-Aggregator/internal/auth"
	"github.com/Insider-Ekin-Aslan/Web-Application-RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handlerCreateUser(writer http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(request.Body)

	params := parameters{}
	decoderError := decoder.Decode(&params)

	if decoderError != nil {
		respondError(writer, 400, fmt.Sprintf("JSON parsing error. ERROR: %v", decoderError))
		return
	}

	user, userError := config.Database.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if userError != nil {
		respondError(writer, 400, fmt.Sprintf("Couldn't create user. ERROR: %v", userError))
		return
	}

	respondJSON(writer, 201, databaseUserToUser(user))
}

func (config *apiConfig) handlerGetUser(writer http.ResponseWriter, request *http.Request) {
	apiKey, apiKeyError := auth.GetAPIKey(request.Header)

	if apiKeyError != nil {
		respondError(writer, 403, fmt.Sprintf("Authentication error. ERROR: %v", apiKeyError))
		return
	}

	user, userError := config.Database.GetUserByAPIKey(request.Context(), apiKey)

	if userError != nil {
		respondError(writer, 400, fmt.Sprintf("Couldn't get user. ERROR: %v", userError))
		return
	}

	respondJSON(writer, 200, databaseUserToUser(user))
}
