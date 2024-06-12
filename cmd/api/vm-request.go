package main

import (
	"Jay179-sudo/backend-api-infra/internal/data"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func (app *application) requestVM(w http.ResponseWriter, r *http.Request) {
	// Set for Testing, change during production
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// extract data from JSON
	VMRequest := data.VMRequest{}
	err := json.NewDecoder(r.Body).Decode(&VMRequest)
	if err != nil {
		app.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	// // validate email using regex

	// Object will enclose and its name/label will be the email provided - (RAM|STORAGE|SERVICE)
	// 1. RAM ------> 	Request + Limit
	// 2. Storage --> 	PV, PVC, Storage Class
	// 3. Service --> 	NodePort

	// send request to the kubernetes API
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Could not authenticate with the API. Error: %v", err.Error())
		http.Error(w, "first The server encountered an error while processing your data", http.StatusInternalServerError)
	}
	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Could not instantiate clientset. Error: %v", err.Error())
		http.Error(w, "second-lastThe server encountered an error while processing your data", http.StatusInternalServerError)
		return
	}
	reducedEmail := ""
	for _, element := range VMRequest.Email {
		if element == '@' {
			break
		} else if element == '_' {
			continue
		}
		reducedEmail += string(element)
	}
	reducedEmail += "pod"
	secondsLeft := time.Until(VMRequest.Spec.ExpiryTime).Seconds()
	if secondsLeft > float64(MaxAllotedTime) || secondsLeft < 0 {
		secondsLeft = float64(MaxAllotedTime)
	}
	allotedSeconds := strconv.Itoa(int(secondsLeft))
	log.Println(allotedSeconds)

	err = app.WriteJson(w, http.StatusOK, envelope{"status": "Request queued!"}, nil)
	if err != nil {
		app.WriteJson(w, http.StatusInternalServerError, envelope{"status": "Internal Server Error"}, nil)
		log.Fatalf("Could not create the Pod %v", err.Error())
		return
	}
}
