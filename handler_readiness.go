package main

import "net/http"

func handlerReadiness(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, 200, struct{}{})
}
