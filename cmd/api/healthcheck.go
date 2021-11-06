package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	healthInfo := jsonWrapper{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version": version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, healthInfo, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}