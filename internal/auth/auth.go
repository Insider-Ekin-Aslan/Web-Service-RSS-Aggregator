package auth

import (
	"errors"
	"net/http"
	"strings"
)

// * Extracts an API Key from the headers of an HTTP Request.
func GetAPIKey(headers http.Header) (string, error) {
	authorisation := headers.Get("Authorization")

	if authorisation == "" {
		return "", errors.New("No authentication header found.")
	}

	array := strings.Split(authorisation, " ")

	if len(array) != 2 || array[0] != "API-Key" {
		return "", errors.New("Malformed authentication header.")
	}

	return array[1], nil
}
