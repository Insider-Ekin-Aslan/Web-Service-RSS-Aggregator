package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (config *apiConfig) handlerCreateUser(writer http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name string `name`
	}

	decoder := json.NewDecoder(request.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondError(writer, 400, fmt.Sprintf("JSON parsing error. ERROR:", err))
		return
	}

	respondJSON(writer, 200, struct{}{})
}
