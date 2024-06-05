package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]interface{}

func dataToJson(data envelope) ([]byte, error) {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return nil, err
	}
	js = append(js, '\n')
	return js, nil
}

func (app *application) WriteJson(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := dataToJson(data)
	if err != nil {
		return err
	}
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
