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

	router.HandlerFunc(http.MethodPost, "/create", app.createShortUrlHandler)
	router.HandlerFunc(http.MethodGet, "/{:id}", app.expandShortUrlHandler)

	router.HandlerFunc(http.MethodGet, "/stats", app.showStatsHandler)

	router.HandlerFunc(http.MethodDelete, "/{:id}", app.deleteShortUrlHandler)

	return router
}