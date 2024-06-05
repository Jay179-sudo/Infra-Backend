package main

import (
	"fmt"
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, status int, message interface{}) {
	env := envelope{"error": message}
	// if error, return 500 Internal Server Error
	err := app.WriteJson(w, status, env, nil)
	if err != nil {
		app.logger.Printf(err.Error())
		w.WriteHeader(500)
	}
}

// implement a server error response method too
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	mp1 := make(map[string]string)
	mp1["message"] = "the requested resource could not be found"
	app.errorResponse(w, http.StatusNotFound, mp1)
}
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	mp1 := make(map[string]string)
	mp1["message"] = fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, http.StatusMethodNotAllowed, mp1)
}
