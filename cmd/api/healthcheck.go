package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Set for Testing, change during production
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
