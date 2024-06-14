package main

import (
	"Jay179-sudo/backend-api-infra/internal/data"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func (app *application) requestVM(w http.ResponseWriter, r *http.Request) {

	// extract data from JSON
	VMRequest := data.VMRequest{}
	err := json.NewDecoder(r.Body).Decode(&VMRequest)
	if err != nil {
		app.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	// Object will enclose and its name/label will be the email provided - (RAM|STORAGE|SERVICE)
	// 1. RAM ------> 	Request + Limit
	// 2. Storage --> 	PV, PVC, Storage Class
	// 3. Service --> 	NodePort

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
	for _, element := range VMRequest.Email {
		if element == '@' {
			break
		} else if element == '_' {
			continue
		}
		reducedEmail += string(element)
	}
	secondsLeft := time.Until(VMRequest.Spec.ExpiryTime).Seconds()
	if secondsLeft > float64(MaxAllotedTime) || secondsLeft < 0 {
		secondsLeft = float64(MaxAllotedTime)
	}
	allotedSeconds := strconv.Itoa(int(secondsLeft))
	log.Println(allotedSeconds)
	userRequestSpec := data.UserRequestSpec{
		Email:     VMRequest.Email,
		RAM:       VMRequest.Spec.RAM,
		CPU:       2,
		PublicKey: VMRequest.Spec.PublicKey,
	}
	userRequest := data.UserRequest{
		TypeMeta: metav1.TypeMeta{
			Kind:       "UserRequest",
			APIVersion: "request.jaypd.github.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      reducedEmail + "-userrequest",
			Namespace: "default",
		},
		Spec: userRequestSpec,
	}
	userReqBody, err := json.Marshal(userRequest)
	if err != nil {
		app.WriteJson(w, http.StatusInternalServerError, envelope{"status": "Internal Server Error"}, nil)
		log.Printf("Could not create the resource %v:", err.Error())
		return
	}
	response := client.RESTClient().
		Post().
		AbsPath("apis/request.jaypd.github.com/v1/namespaces/default/userrequests").
		Body(userReqBody).
		Do(context.TODO())
	if err = response.Error(); err != nil {
		app.WriteJson(w, http.StatusInternalServerError, envelope{"status": "Internal Server Error"}, nil)
		log.Printf("Could not create the resource %v:", err.Error())
		return
	}
	err = app.WriteJson(w, http.StatusOK, envelope{"status": "Request queued!"}, nil)
	if err != nil {
		app.WriteJson(w, http.StatusInternalServerError, envelope{"status": "Internal Server Error"}, nil)
		log.Printf("Could not create the resource %v:", err.Error())
		return
	}
}
