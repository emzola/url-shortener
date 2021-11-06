package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/emzola/url-shortener/base62"
	"github.com/emzola/url-shortener/models"
	"github.com/emzola/url-shortener/validator"
	"github.com/julienschmidt/httprouter"
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

func (app *application) createShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		URL string	`json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	webUrl := &models.Url{
		URL: input.URL,
	}

	v := validator.IsUrl(webUrl.URL)
	if !v {
		app.failedValidationResponse(w, r, err)
		return
	}
	
	err = app.model.Insert(webUrl)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var id = base62.Encode(webUrl.ID)

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/shorturl/%s", id))

	err = app.writeJSON(w, http.StatusCreated, jsonWrapper{"short_url": id}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) expandShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err :=  base62.Decode(params.ByName("id"))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	url, err := app.model.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if r.Method == "GET" {
    http.Redirect(w, r, url.URL, http.StatusMovedPermanently)
	}
}

func (app *application) deleteShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete short url")
}