package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthCheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/shorturl", app.createShortUrlHandler)
	router.HandlerFunc(http.MethodGet, "/v1/shorturl/:id", app.expandShortUrlHandler)

	router.HandlerFunc(http.MethodDelete, "/v1/shorturl/:id", app.deleteShortUrlHandler)

	return router
}