package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emzola/url-shortener/base62"
	"github.com/emzola/url-shortener/models"
	"github.com/emzola/url-shortener/validator"
	"github.com/julienschmidt/httprouter"
)

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
	headers.Set("Location", fmt.Sprintf("/%v", id))

	err = app.writeJSON(w, http.StatusCreated, jsonWrapper{"shortUrl": fmt.Sprintf("localhost:%d/%v", app.config.port, id)}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) expandShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id := params.ByName("id")

	fmt.Fprintf(w, "show details of url %s\n", id)
}

func (app *application) showStatsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "show stats handler")
}

func (app *application) deleteShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete short url")
}