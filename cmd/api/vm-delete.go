package main

import (
	"Jay179-sudo/backend-api-infra/internal/data"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func (app *application) deleteVM(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete VM request triggered")
	DeleteRequest := data.DeleteRequest{}
	err := json.NewDecoder(r.Body).Decode(&DeleteRequest)
	if err != nil {
		log.Printf("Could not process request")
		app.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	// authenticate with the server
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("Could not authenticate with the API. Error: %v", err.Error())
		http.Error(w, "The server encountered an error while processing your data", http.StatusInternalServerError)
		return
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Could not instantiate clientset. Error: %v", err.Error())
		http.Error(w, "The server encountered an error while processing your data", http.StatusInternalServerError)
		return
	}
	reducedEmail := ""
	for _, element := range DeleteRequest.Email {
		if element == '@' {
			break
		} else if element == '_' {
			continue
		}
		reducedEmail += string(element)
	}
	response := client.RESTClient().
		Delete().
		AbsPath("apis/request.jaypd.github.com/v1/namespaces/default/userrequests/" + reducedEmail + "-userrequest").
		Do(context.TODO())
	if err = response.Error(); err != nil {
		app.WriteJson(w, http.StatusInternalServerError, envelope{"status": "Internal Server Error"}, nil)
		log.Printf("Could not create the resource %v:", err.Error())
		return
	}
	err = app.WriteJson(w, http.StatusOK, envelope{"status": "Resource Deleted!"}, nil)
	if err != nil {
		app.WriteJson(w, http.StatusInternalServerError, envelope{"status": "Internal Server Error"}, nil)
		log.Printf("Could not create the resource %v:", err.Error())
		return
	}
}
