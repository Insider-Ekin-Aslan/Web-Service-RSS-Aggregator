package main

import (
	"fmt"
	"net/http"

	"github.com/Insider-Ekin-Aslan/Web-Application-RSS-Aggregator/internal/auth"
	"github.com/Insider-Ekin-Aslan/Web-Application-RSS-Aggregator/internal/database"
)

type authenticatedHandler func(http.ResponseWriter, *http.Request, database.User)

func (config *apiConfig) middlewareAuth(handler authenticatedHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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

		handler(writer, request, user)
	}
}
