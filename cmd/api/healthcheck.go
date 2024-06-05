package main

import "net/http"

func (app *application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	content := map[string]string{
		"status":      "Available",
		"environment": app.config.env,
	}
	err := app.WriteJson(w, http.StatusOK, envelope{"status": content}, nil)
	if err != nil {
		app.logger.Println(err.Error())
		http.Error(w, "The server encountered an error while processing your data", http.StatusInternalServerError)
	}

}
