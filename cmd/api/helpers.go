package main

import (
	"encoding/json"
	"net/http"
)

// Key wrapper for json objects
type jsonWrapper map[string]interface{}

// Helper method for encoding json
func (app *application) writeJSON(w http.ResponseWriter, status int, data jsonWrapper, headers http.Header) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	jsonData = append(jsonData, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)

	return nil
}