package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/api/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/api/request", app.requestVM)
	router.HandlerFunc(http.MethodPost, "/api/delete", app.deleteVM)
	return app.enableCORS(router)
}
