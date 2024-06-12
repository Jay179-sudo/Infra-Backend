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

func (app *application) readJson(r *http.Request, input interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return err
	}
	return nil
}

// func connectDB(connectionString string) (*mongo.Client, error) {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// 	return client, err
// }

// func (app *application) createIndex() error {
// 	indexModel := mongo.IndexModel{
// 		Keys:    bson.D{{"Email", 1}},
// 		Options: options.Index().SetUnique(true),
// 	}
// 	coll := app.mongo.Database("User-Requests").Collection("VM")
// 	_, err := coll.Indexes().CreateOne(context.TODO(), indexModel)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
