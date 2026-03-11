package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondJSON(writer http.ResponseWriter, code int, payload interface{}) {
	data, jsonError := json.Marshal(payload)

	if jsonError != nil {
		log.Println("Failed to marshal JSON response.")
		log.Println(payload)

		writer.WriteHeader(500)

		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(data)
}

func respondError(writer http.ResponseWriter, code int, message string) {
	if code >= 500 {
		log.Println("Responding with 5XX error")
		log.Println(message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondJSON(writer, code, errorResponse{Error: message})
}
